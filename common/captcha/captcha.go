package captcha

import (
	"context"
	"sync"

	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}

// 确保执行一次
var once sync.Once

// 内部使用
var captcha *Captcha

// NewCaptcha 单例模式获取
func NewCaptcha(ctx context.Context, client *redis.Redis) *Captcha {
	once.Do(func() {
		// 初始化captcha对象
		captcha = &Captcha{}

		// 获得store： 存储信息
		store := &RedisStore{
			RedisClient: client,
			Ctx:         ctx,
		}

		// 获得driver：产生验证码信息
		// 后面要加上从配置信息中获得
		driver := base64Captcha.NewDriverDigit(
			80,  // width
			240, // height
			6,
			0.7,
			80,
		)

		captcha.Base64Captcha = base64Captcha.NewCaptcha(driver, store)
	})

	return captcha
}

// GenerateCaptcha: 生成验证码
func (c *Captcha) GenerateCaptcha() (id, b64s, answer string, err error) {
	return c.Base64Captcha.Generate()
}

func (c *Captcha) VerifyCaptcha(id, answer string) (match bool) {
	return c.Base64Captcha.Store.Verify(id, answer, true)
}
