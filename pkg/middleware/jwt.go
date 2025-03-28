package middleware

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rongpengju/gin-template/configs"
)

// token过期时间, 如果需要按环境定制化, 请写到配置文件中
const tokenExpireDuration = time.Hour * 8

var customSecret = []byte("gin-template")

// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
type CustomClaims struct {
	Uuid                 string `json:"uuid"` // 用户 uuid
	jwt.RegisteredClaims        // 内嵌标准的声明
}

// GenerateJwtToken 生成JWT
func GenerateJwtToken(uuid string) (string, error) {
	claims := CustomClaims{
		Uuid: uuid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpireDuration)),
			Issuer:    configs.Conf.App.Name,
		},
	}

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(customSecret)
}

// ParseJwtToken 解析JWT
func ParseJwtToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return customSecret, nil
	})
	if err != nil {
		return nil, err
	}

	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}

	return nil, fmt.Errorf("token is invalid")
}
