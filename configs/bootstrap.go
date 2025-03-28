package configs

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

const (
	configDir    = "configs/"
	configPrefix = "gin-template"
	configType   = "yaml"
	osEnvKey     = "ENV"
)

func init() {
	// 从环境配置中读取 ENV 的值
	env := os.Getenv(osEnvKey)
	if env == "" {
		env = gin.DebugMode
	}

	// 根据文件地址读取文件内容
	configFilePath := fmt.Sprintf("%s%s-%s.%s", configDir, configPrefix, env, configType)
	fileContent, err := os.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}

	// 文件内容反序列化到结构体上
	if err := yaml.Unmarshal(fileContent, &Conf); err != nil {
		panic(err)
	}

	// 如果启动的是 release 环境则作出警告提示
	if env == gin.ReleaseMode {
		fmt.Printf("%s当前启动环境为【release】请谨慎操作数据！%s\n", "\033[31m", "\033[0m")
	}
}
