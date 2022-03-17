package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	TokenExpireDuration = time.Hour * 24 * 3
	Mysecret            = "https://github.com/LUZIWEI960903"
)

var InvalidToken = errors.New("Invalid Token!!")

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
type MyClaims struct {
	Username           string `json:"username"` // 自定义签名字段
	UserID             int64  `json:"user_id"`  // 自定义签名字段
	jwt.StandardClaims        // jwt包自带的jwt.StandardClaims只包含了官方字段
}

// GenToken 生成JWT
func GenToken(username string, userid int64) (string, error) {
	// 构造一个自定义的声明
	c := &MyClaims{
		Username: username,
		UserID:   userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "Ziwei.Lu",
		},
	}

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(Mysecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Mysecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, InvalidToken
}
