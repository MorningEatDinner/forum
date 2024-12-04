package mq

// SendPhoneCodeMessage 发送手机验证码消息
type SendPhoneCodeMessage struct {
	Phone string
	Code  string
}

// SendEmailCodeMessage 发送邮箱验证码消息
type SendEmailCodeMessage struct {
	Email string
	Code  string
}

type DeletePostMessage struct {
	PostId int64
}

type UpdatePostScore struct {
	PostId int64
	Score  int64
	Up     bool
	Down   bool
}
