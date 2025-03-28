package library

import (
	"github.com/rongpengju/gin-template/library/mini_program"
	"github.com/rongpengju/gin-template/library/wechat_work"
)

func InitAllLibrary() {
	mini_program.InitMiniMiniProgramLibrary()
	wechat_work.InitWechatWorkLibrary()
}
