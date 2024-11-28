package mail

import "context"

type Driver interface {
	// 发送验证码
	Send(ctx context.Context, email Email) bool
}
