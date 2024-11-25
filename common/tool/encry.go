package tool

import (
	"crypto/md5"
	"encoding/hex"
)

const secret = "forum.xiaorui.com"

// 使用md5加密密码现在被认为是不安全的了， 因为可以通过如暴力破解的方式来破解
func EncryptPassword(oPasssword string) string {
	h := md5.New()                                       // 创建hash.Hash对象
	h.Write([]byte(secret))                              //  写入颜值
	return hex.EncodeToString(h.Sum([]byte(oPasssword))) // 对于数据进行加密
}

func CheckPasswordHash(password, hashedPassword string) bool {
	h := md5.New()
	h.Write([]byte(secret))
	hashedInput := hex.EncodeToString(h.Sum([]byte(password)))
	return hashedInput == hashedPassword
}
