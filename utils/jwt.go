package utils

import (
	"blog/configs"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	UserID int `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(userID int) (string, error) {
	cfg, err := configs.GetConfig()
	if err != nil {
		fmt.Println("Error while getting config", err)
		return "", err
	}

	expirationTime := time.Now().Add(24 * time.Hour * 7)

	// 创建一个 Claims 对象并设置其字段值
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// 使用声明的算法创建一个新的 JWT 对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥进行签名并获取字符串格式的令牌
	tokenString, err := token.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	cfg, err := configs.GetConfig()
	if err != nil {
		return nil, err
	}

	// 解析 JWT 字符串
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	// 验证 token 是否有效
	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	// 将解析后的 Claims 类型断言为我们定义的 Claims 结构体类型
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, jwt.ErrInvalidKey
	}

	return claims, nil
}
