package adapter

import (
	"crypto/rand"
	"encoding/hex"
	"login-server/internal/models"
	"login-server/internal/types"
	"login-server/utils"
	"strings"
)

const DefaultAvatar = "http://example.com/avatar.jpg"

func BuildInsertUser(req *types.RegisterReq, ip string) (*models.User, error) {
	password, err := utils.SecretPassword(req.Password)
	if err != nil {
		return nil, err
	}
	ipInfo := []string{ip}
	strings.Join(ipInfo, ",")
	user := models.NewUserBuilder().
		SetName(newUserName()).
		SetAvatar(DefaultAvatar).
		SetSex(1).
		SetStatus(0).
		SetIPInfo(ipInfo).
		SetEmail(req.Email).
		SetPassword(password).
		Build()
	return &user, nil
}

func newUserName() string {
	randomBytes := make([]byte, 8) // 生成8个随机字节
	_, err := rand.Read(randomBytes)
	if err != nil {
		// 处理错误
		return "用户" // 出现错误时返回默认值
	}
	randomString := hex.EncodeToString(randomBytes) // 将随机字节转换为十六进制字符串
	return "用户" + randomString
}
