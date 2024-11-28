package mail

import (
	"context"
	"fmt"
	"net/smtp"

	emailPKG "github.com/jordan-wright/email"
	"github.com/zeromicro/go-zero/core/logx"
)

type Config struct {
	*SmptConfig `json:"smtp"`
	*FromConfig `json:"from"`
}

type SmptConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type FromConfig struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type SMTP struct {
	config *SmptConfig
}

// Send: 发送邮箱验证码
func (s *SMTP) Send(ctx context.Context, email Email) bool {
	e := emailPKG.NewEmail()

	e.From = fmt.Sprintf("%v <%v>", email.From.Name, email.From.Address)
	e.To = email.To
	e.Bcc = email.Bcc
	e.Cc = email.Cc
	e.Subject = email.Subject
	e.Text = email.Text
	e.HTML = email.HTML
	err := e.Send(
		fmt.Sprintf("%v:%v", s.config.Host, s.config.Port),
		smtp.PlainAuth(
			"",
			s.config.Username,
			s.config.Password,
			s.config.Host,
		),
	)
	if err != nil {
		logx.WithContext(ctx).Errorf("send email failed, err: %v", err.Error())
		return false
	}

	return true
}
