package sms

type SMSConfig struct {
	// AccessKeyID     string `json:"access_key_id"`
	// AccessKeySecret string `json:"access_key_secret"`
	// 上面从环境变量中获取
	SignName     string `json:"sign_name"`
	TemplateCode string `json:"template_code"`
}
