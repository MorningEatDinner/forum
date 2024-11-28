package svc

import (
	"forum/app/user/mq/internal/config"
	"forum/common/mail"
	"forum/common/sms"
)

type ServiceContext struct {
	Config config.Config

	SMSClient  *sms.Aliyun
	MailClient *mail.Mailer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		SMSClient:  sms.NewSmsClient(c.Sms),
		MailClient: mail.NewMailer(c.MailConf),
	}
}
