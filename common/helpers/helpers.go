package helpers

import (
	"crypto/rand"
	"fmt"
	"io"
	mathrand "math/rand"
	"time"
)

// FirstElement 安全地获取 args[0]，避免 panic: runtime error: index out of range
func FirstElement(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return ""
}

// RandomString: 生成长度为length的随机字符串
func RandomString(length int) string {
	mathrand.New(mathrand.NewSource(time.Now().UnixNano()))
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length) // rune是32位
	for i := range b {
		b[i] = letters[mathrand.Intn(len(letters))]
	}
	return string(b)
}

// RandomNumber 生成长度为 length 随机数字字符串
func RandomNumber(length int) string {
	table := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, length)
	n, err := io.ReadAtLeast(rand.Reader, b, length)
	if n != length {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

// GenerateRandomCode: 生成随机验证码
func GenerateRandomCode() string {
	code := fmt.Sprintf("%06v", mathrand.New(mathrand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	return code
}
