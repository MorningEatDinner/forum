package logic

import (
	"context"
	"fmt"
	"forum/app/mqueue/cmd/job/internal/svc"
	"forum/app/post/rpc/postservice"
	"forum/app/user/rpc/userservice"
	"forum/common/mail"
	"forum/common/xerr"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"github.com/sourcegraph/conc"
	"github.com/sourcegraph/conc/pool"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/time/rate"
)

const (
	TopN      = 10
	BatchSize = 100
)

type HotPostPushingHandler struct {
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	limiter *rate.Limiter
}

func NewHotPostPushingHandler(svcCtx *svc.ServiceContext) *HotPostPushingHandler {
	return &HotPostPushingHandler{
		ctx:     context.Background(),
		svcCtx:  svcCtx,
		limiter: rate.NewLimiter(rate.Limit(20), 10), // 限制最大并发10个， 限制每秒最多20个请求
	}
}

// 这里处理延迟的那个任务
func (l *HotPostPushingHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	// 1. 获取前十的热点帖子
	postList, err := l.svcCtx.PostRpc.GetPostList(l.ctx, &postservice.GetPostListRequest{
		Page:     1,
		PageSize: TopN,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("failed to get post list: %v", err)
		return errors.Wrapf(xerr.NewErrMsg("failed to get post list"), "failed to get post list")
	}

	// 创建用户信息通道
	feeder := make(chan *userservice.UserInfo, 100)

	// 启动后台消费者协程
	var wg conc.WaitGroup
	wg.Go(func() {
		p := pool.New().WithMaxGoroutines(10)
		for user := range feeder {
			user := user
			p.Go(func() {
				if user.Email == "" {
					logx.WithContext(l.ctx).Infof("user %d has no email", user.Username)
					return
				}

				// 使用限流器
				if err := l.limiter.Wait(l.ctx); err != nil {
					logx.WithContext(l.ctx).Errorf("rate limit error: %v", err)
					return
				}

				email := mail.Email{
					From: mail.From{
						Address: l.svcCtx.Config.MailConf.FromConfig.Address,
						Name:    l.svcCtx.Config.MailConf.FromConfig.Name,
					},
					To:      []string{user.Email},
					Subject: "Forum 论坛每周热点帖子",
					HTML:    []byte(fmt.Sprintf(mail.TemplateHTMLWeeklyHotPosts, user.Username, mail.GenerateAllPostsHTML(postList.Posts))),
				}

				if !l.svcCtx.MailClient.Send(l.ctx, email) {
					logx.WithContext(l.ctx).Errorf("failed to send email to user: %v", err)
				}
			})
		}
		p.Wait() // 等待所有邮件发送完成
	})
	// 生产者：批量获取用户并写入通道
	var lastUid int64
	for {
		users, err := l.svcCtx.UserRpc.GetUserList(l.ctx, &userservice.GetUserListRequest{
			LastUserId: lastUid,
			BatchSize:  BatchSize,
		})
		if err != nil {
			logx.WithContext(l.ctx).Errorf("failed to get user list: %v", err)
			close(feeder) // 发生错误时关闭通道
			wg.Wait()     // 等待消费者处理完剩余数据
			return errors.Wrapf(xerr.NewErrMsg("failed to get user list"), "failed to get user list")
		}

		// 如果没有获取到任何用户，说明已经处理完所有用户
		if len(users.Users) == 0 {
			logx.WithContext(l.ctx).Info("no more users to process")
			break
		}

		// 将用户信息写入通道
		for _, user := range users.Users {
			feeder <- user
		}

		// 如果是最后一批数据，结束循环
		if len(users.Users) < BatchSize {
			break
		}

		lastUid = users.LastUserId
	}

	// 所有用户数据获取完成，关闭通道
	close(feeder)

	// 等待消费者处理完所有数据
	wg.Wait()

	return nil
}
