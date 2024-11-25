package sms

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
)

// 阿里云实现发送短信接口
type Aliyun struct {
	client       *dysmsapi20170525.Client
	SignName     string
	TemplateCode string
}

func (s *Aliyun) Send(ctx context.Context, phone, message string) bool {
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(phone),
		SignName:      tea.String(s.SignName),
		TemplateCode:  tea.String(s.TemplateCode),
		TemplateParam: tea.String(fmt.Sprintf("{\"code\":\"%s\"}", message)),
	}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, err := s.client.SendSmsWithOptions(sendSmsRequest, &util.RuntimeOptions{})
		if err != nil {
			logx.WithContext(ctx).Error("短信[阿里云]\", \"数据发送失败", zap.Error(err))
			return err
		}

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 错误 message
		fmt.Println(tea.StringValue(error.Message))
		// 诊断地址
		var data interface{}
		d := json.NewDecoder(strings.NewReader(tea.StringValue(error.Data)))
		d.Decode(&data)
		if m, ok := data.(map[string]interface{}); ok {
			recommend, _ := m["Recommend"]
			fmt.Println(recommend)
		}
		_, err := util.AssertAsString(error.Message)
		if err != nil {
			logx.WithContext(ctx).Error("短信[阿里云]\", \"数据发送失败", zap.Error(err))
			return false
		}
	}
	return true
}

// CreateClient
/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func createClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func NewSmsClient(config SMSConfig) *Aliyun {
	cli, err := createClient(&config.AccessKeyID, &config.AccessKeySecret)
	if err != nil {
		logx.Errorf("初始化阿里云短信客户端失败: %v", err)
		panic(err) // 或者直接 os.Exit(1)
	}
	return &Aliyun{
		client:       cli,
		SignName:     config.SignName,
		TemplateCode: config.TemplateCode,
	}
}
