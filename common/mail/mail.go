package mail

import (
	"context"
	"sync"
)

const (
	TemplateHTML = `
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
