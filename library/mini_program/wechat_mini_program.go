package mini_program

import (
	"sync"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"

	"github.com/rongpengju/gin-template/configs"
)

var (
	miniProgramLibrary *miniProgram.MiniProgram
	once               sync.Once
)

// InitMiniMiniProgramLibrary 初始化微信小程序库
func InitMiniMiniProgramLibrary() {
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
		miniProgramLibrary, err = miniProgram.NewMiniProgram(&miniProgram.UserConfig{
			AppID:  "APP_ID",
			Secret: "APP_SECRET",
			Cache:  cache,
		})
		if err != nil {
			panic(err)
		}
	})
}
