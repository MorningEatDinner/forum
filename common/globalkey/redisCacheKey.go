package globalkey

const (
	KeyPrefix = "forum:"
	// CacheUserTokenKey /** 用户登陆的token
	CacheUserTokenKey = "user_token:%v"
	// 验证码
	CaptchaKey   = "captcha:%v"
	PhoneCodeKey = "phone_code:%v"
)

// 给key加上前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}

func GetRedisKey(key string) string {
	return getRedisKey(key)
}
