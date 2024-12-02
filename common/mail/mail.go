package mail

import (
	"context"
	"fmt"
	"forum/app/post/rpc/postservice"
	"strings"
	"sync"
)

const (
	TemplateHTMLCode = `
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="utf-8">
		<title>验证码</title>
	</head>
	<body>
		<div style="max-width: 600px; margin: 0 auto; padding: 20px; font-family: Arial, sans-serif;">
			<h2 style="color: #333;">验证码</h2>
			<p>您的验证码是：</p>
			<div style="background: #f4f4f4; padding: 10px; margin: 20px 0; text-align: center;">
				<span style="font-size: 24px; font-weight: bold; letter-spacing: 5px;">%s</span>
			</div>
			<p style="color: #666; font-size: 14px;">验证码 10 分钟内有效，请勿泄露给他人。</p>
		</div>
	</body>
	</html>`

	TemplateText = `验证码：%s
	验证码 10 分钟内有效，请勿泄露给他人。`

	TemplateHTMLProfileUpdate = `
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="utf-8">
		<title>个人信息更新提醒</title>
	</head>
	<body>
		<div style="max-width: 600px; margin: 0 auto; padding: 20px; font-family: Arial, sans-serif;">
			<h2 style="color: #333;">温馨提示</h2>
			<p>亲爱的 %s：</p>
			<div style="background: #f4f4f4; padding: 15px; margin: 20px 0; line-height: 1.6;">
				<p>我们注意到您已注册成为我们的用户一天了，但您的个人信息还未完善。</p>
				<p>完整的个人信息有助于我们为您提供更好的服务体验。</p>
				<p>请尽快登录系统完善您的个人信息。</p>
			</div>
			<p style="color: #666; font-size: 14px;">如果您已经更新了信息，请忽略此邮件。</p>
			<p style="color: #666; font-size: 14px;">感谢您的支持！</p>
		</div>
	</body>
	</html>`

	// TemplateHTMLWeeklyHotPosts 定义邮件模板
	TemplateHTMLWeeklyHotPosts = `
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="utf-8">
		<title>本周热点帖子推荐</title>
	</head>
	<body>
		<div style="max-width: 600px; margin: 0 auto; padding: 20px; font-family: Arial, sans-serif;">
			<h2 style="color: #333;">本周热点精选</h2>
			<p>亲爱的 %s：</p>
			<div style="background: #f4f4f4; padding: 15px; margin: 20px 0; line-height: 1.6;">
				<p>以下是本周最受欢迎的热门帖子，希望对您有帮助：</p>
				%s
			</div>
			<p style="color: #666; font-size: 14px;">更多精彩内容，欢迎随时访问我们的网站查看。</p>
			<p style="color: #666; font-size: 14px;">祝您学习愉快！</p>
			<div style="margin-top: 20px; padding-top: 20px; border-top: 1px solid #eee; font-size: 12px; color: #999;">
				<p>如果您不想继续接收此类邮件，可以点击此处取消订阅。</p>
			</div>
		</div>
	</body>
	</html>`
)

type From struct {
	Address string // email地址
	Name    string // 名字
}

type Email struct {
	From    From     // 发件人
	To      []string // 收件人
	Bcc     []string // 密送
	Cc      []string // 抄送
	Subject string   // 主题
	Text    []byte   // 纯文本
	HTML    []byte
}

type Mailer struct {
	Driver Driver
}

var once sync.Once
var mailer *Mailer

func NewMailer(config Config) *Mailer {
	once.Do(func() {
		mailer = &Mailer{
			Driver: &SMTP{
				config: config.SmptConfig,
			},
		}
	})

	return mailer
}

func (m *Mailer) Send(ctx context.Context, email Email) bool {
	return m.Driver.Send(ctx, email)
}

// 生成所有帖子的HTML内容
func GenerateAllPostsHTML(posts []*postservice.Post) string {
	var postsHTML strings.Builder
	for _, post := range posts {
		postsHTML.WriteString(generatePostHTML(post))
	}
	return postsHTML.String()
}

// 生成单个帖子的HTML内容
func generatePostHTML(post *postservice.Post) string {
	return fmt.Sprintf(`
        <div style="background: white; padding: 12px; margin: 10px 0; border-radius: 5px; box-shadow: 0 1px 3px rgba(0,0,0,0.1);">
            <h3 style="color: #2c5282; margin: 0 0 8px 0; font-size: 16px;">%s</h3>
            <p style="color: #666; margin: 0; font-size: 14px;">%s</p>
        </div>`, post.Title, post.Content)
}
