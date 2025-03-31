package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/MrHat365/odin-go/odin_api/client"
	"github.com/MrHat365/odin-go/odin_api/models"
)

// AuthIdentity 身份验证和身份注册请求
// identity参数需要实现GetPublicKey()和Sign()方法
type Identity interface {
	GetPublicKey() []byte
	Sign(message []byte) ([]byte, error)
}

// AuthIdentity 身份验证和身份注册请求
// 与支持ED25519的身份验证一起使用来生成签名并获取授权令牌
func AuthIdentity(identity Identity) (*models.AuthToken, error) {
	c := client.NewClient()

	// 获取当前时间戳
	now := time.Now().UnixMilli()
	timestamp := fmt.Sprintf("%d", now)

	// 签名时间戳
	signature, err := identity.Sign([]byte(timestamp))
	if err != nil {
		return nil, fmt.Errorf("签名生成失败: %w", err)
	}

	// 创建授权请求
	authReq := models.AuthRequest{
		PublicKey: base64.StdEncoding.EncodeToString(identity.GetPublicKey()),
		Timestamp: timestamp,
		Signature: base64.StdEncoding.EncodeToString(signature),
		Referrer:  "zg8khi8rz0",
	}

	// 发送请求
	resp, err := c.Post("/auth", authReq)
	if err != nil {
		return nil, fmt.Errorf("身份验证请求失败: %w", err)
	}

	// 解析响应
	var authToken models.AuthToken
	if err := json.Unmarshal(resp, &authToken); err != nil {
		return nil, fmt.Errorf("解析授权令牌失败: %w", err)
	}

	return &authToken, nil
}
