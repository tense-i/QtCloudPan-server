package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSecret = []byte("QtCloudPanBytensei")

// 自定义的声明 (claims)
type CustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// 生成JWT
func GenerateToken(username string, duration time.Duration) (string, error) {
	// 创建声明
	claims := CustomClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)), // 设置过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),               // 签发时间
		},
	}

	// 创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名并获得完整的编码后的token字符串
	return token.SignedString(jwtSecret)
}

// 验证JWT
func ValidateToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 确保使用的是HMAC方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	// 解析错误或token无效
	if err != nil {
		return nil, err
	}

	// 验证token的声明 (claims)
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}

// 解析JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	// 解析错误或token无效
	if err != nil {
		return nil, err
	}

	// 验证token的声明 (claims)
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}

//
