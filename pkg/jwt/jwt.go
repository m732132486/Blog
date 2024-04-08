package jwt

import (
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var mySecret = []byte("secret")

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"user"`
	jwt.StandardClaims
}

// GenerateJWT 生成jwt
func GenerateJWT(username string, userID int64) (string, error) {
	//创建一个自己的声明
	c := MyClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(
				time.Duration(viper.GetInt("auth.jwt_expire")) * time.Second).Unix(), //到期时间
			Issuer: "dd", //签发人
		},
	}
	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(mySecret)
}

// ParseToken 解析函数
func ParseToken(tokenString string) (*MyClaims, error) {
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}

// ParseTokenID 根据token解析用户id
func ParseTokenID(tokenString string) (int64, error) {
	// 移除前缀 "Bearer "
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	// 创建 MyClaims 结构体实例
	mc := new(MyClaims)

	// 解析 JWT 令牌，并将声明解析到 MyClaims 结构体中
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil // 这里需要提供用于验证签名的密钥
	})
	if err != nil {
		return 0, err
	}

	// 验证令牌是否有效
	if token.Valid {
		// 提取用户ID并返回
		return mc.UserID, nil
	}

	return 0, errors.New("invalid token")
}
