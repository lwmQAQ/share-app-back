package utils

import (
	"fmt"
	"user-server/config"

	"github.com/golang-jwt/jwt/v5"
)

type JWTUtil struct {
	config *config.JWTConfig
}

type Claims struct {
	ID uint64 `json:"id"`
	jwt.RegisteredClaims
}

func NewJWTUtil(jwtconfig *config.JWTConfig) *JWTUtil {
	return &JWTUtil{
		config: jwtconfig,
	}
}

// 生成token
func (util *JWTUtil) GenerateToken(id uint64) (string, error) {
	claims := Claims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   "lwm开发者",            // 颁发者
			Subject:  "user-id",           // 主题
			Audience: []string{"lwm平台用户"}, // 接受者
		},
	}

	// 创建一个新的 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(token)
	// 使用密钥签名 token
	return token.SignedString([]byte(util.config.Key))
}

func (util *JWTUtil) ParseToken(tokenString string) (uint64, error) {
	// 解析 token，使用 Claims 结构来存储解析出的 claims
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// 确保 token 的签名方法是我们预期的
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(util.config.Key), nil // 返回用于验证签名的密钥
	})

	if err != nil {
		fmt.Printf("Error parsing token: %v\n", err) // 打印解析错误
		return 0, err                                // 返回解析过程中发生的错误
	}

	// 返回解析出的 claims
	return claims.ID, nil
}
