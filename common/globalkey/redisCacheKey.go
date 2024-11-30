package globalkey

const (
	KeyPrefix = "forum:"
	// CacheUserTokenKey /** 用户登陆的token
	CacheUserTokenKey = "user_token:%v"
	// 验证码
	CaptchaKey   = "captcha:%v"
	PhoneCodeKey = "phone_code:%v"
	EmailCodeKey = "email_code:%v"

	// 更新表示
	UpdatedKey = "user_profile_status:%v"

	// 投票记录
	VoteRecordKey = "post_vote:%v"

	// 帖子
	PostScoreKey     = "post_score"
	PostUpCountKey   = "post_up_count"
	PostDownCountKey = "post_down_count"

	// 社区
	PostCommunityKey = "post_community:%v"
)

// 给key加上前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}

func GetRedisKey(key string) string {
	return getRedisKey(key)
}
