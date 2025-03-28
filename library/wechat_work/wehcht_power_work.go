package wechat_work

import (
	"fmt"
	"sync"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work"
	"github.com/go-resty/resty/v2"

	"github.com/rongpengju/gin-template/configs"
)

const robotKey = "YOUR-ROBOT-KEY"

var (
	wechatWorkLibrary *work.Work
	once              sync.Once
)

// InitWechatWorkLibrary 初始化企业微信对接库
func InitWechatWorkLibrary() {
	var cache kernel.CacheInterface
	if configs.Conf.DataSource.Redis.Addr != "" {
		cache = kernel.NewRedisClient(&kernel.UniversalOptions{
			Addrs:    []string{configs.Conf.DataSource.Redis.Addr},
			Password: configs.Conf.DataSource.Redis.Password,
			DB:       configs.Conf.DataSource.Redis.DB,
		})
	}

	once.Do(func() {
		var err error
		wechatWorkLibrary, err = work.NewWork(&work.UserConfig{
			OAuth: work.OAuth{
				Callback: "https://wecom.artisan-cloud.com/callback",
				Scopes:   nil,
			},
			Cache: cache,
		})
		if err != nil {
			panic(err)
		}
	})
}

// RobotNotification 企业微信机器人通知
func RobotNotification(message string) error {
	payload := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": message,
		},
	}
	resp, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", robotKey))
	if err != nil {
		return fmt.Errorf("企业微信机器人发送通知失败，Error：%v， Response：%v", err.Error(), resp.String())
	}

	return nil
}
