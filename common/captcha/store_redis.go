package captcha

import (
	"context"
	"errors"
	"fmt"
	"forum/common/globalkey"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

/*
	实现captcha的redis store
*/

type RedisStore struct {
	RedisClient *redis.Redis
	Ctx         context.Context
}

// Set： 设置验证码id和答案
func (s *RedisStore) Set(id string, value string) error {
	expireTime := time.Minute * 15 // 后面可以在配置中导入
	expireSeconds := int(expireTime.Seconds())

	key := fmt.Sprintf(globalkey.GetRedisKey(globalkey.CaptchaKey), id)
	if err := s.RedisClient.Setex(key, value, expireSeconds); err != nil {
		logx.WithContext(s.Ctx).Error("存储数据失败", err.Error())
		return errors.New("存储数据失败")
	}

	return nil
}

// Get：获得验证码的值
func (s *RedisStore) Get(id string, clear bool) string {
	key := fmt.Sprintf(globalkey.GetRedisKey(globalkey.CaptchaKey), id)
	fmt.Println("key: ", key)
	val, err := s.RedisClient.Get(key)
	if err != nil {
		return ""
	}
	if clear {
		_, _ = s.RedisClient.Del(key)
	}
	return val
}

// Verify：验证验证码是否正确
func (s *RedisStore) Verify(id, answer string, clear bool) bool {
	realVal := s.Get(id, clear)
	return realVal == answer
}
