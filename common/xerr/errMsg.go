package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[SERVER_COMMON_ERROR] = "服务器开小差啦,稍后再来试一试"
	message[REUQEST_PARAM_ERROR] = "参数错误"
	message[TOKEN_EXPIRE_ERROR] = "token失效，请重新登陆"
	message[TOKEN_GENERATE_ERROR] = "生成token失败"
	message[DB_ERROR] = "数据库繁忙,请稍后再试"
	message[DB_UPDATE_AFFECTED_ZERO_ERROR] = "更新数据影响行数为0"
	message[PERMISSION_DENIED] = "权限不足"

	// 用户模块错误
	message[USER_CAPTCHA_ERROR] = "验证码错误"
	message[USER_NOT_FOUND] = "用户不存在"
	message[USER_PASSWORD_ERROR] = "密码错误"
	message[USER_NAME_EXISTS_ERROR] = "用户名已存在"
	message[USER_PASSWORD_DISMATCH_ERROR] = "两次密码输入不一致"

	// 社区模块错误
	message[COMMUNITY_NAME_EXIST] = "社区名称已存在"
	message[COMMUNITY_NOT_EXIST] = "社区不存在"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
