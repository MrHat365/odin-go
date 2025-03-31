package user

import (
	"encoding/json"
	"fmt"

	"github.com/MrHat365/odin-go/odin_api/client"
	"github.com/MrHat365/odin-go/odin_api/models"
)

// ChangeUsername 使用授权令牌更改用户名
// 需要提供用户名、主体ID和授权令牌
func ChangeUsername(username, principalID, authToken string) (string, error) {
	c := client.NewClient()
	c.SetToken(authToken)

	// 创建表单数据
	formData := map[string]string{
		"username": username,
	}

	// 发送请求
	endpoint := fmt.Sprintf("/user/profile?user=%s", principalID)
	resp, err := c.PostMultipart(endpoint, formData)
	if err != nil {
		return "", fmt.Errorf("更改用户名请求失败: %w", err)
	}

	return string(resp), nil
}

// GetOdinFunUser 获取Odin.fun用户信息
// 根据提供的principal_id获取用户信息
func GetOdinFunUser(principalID string) (*models.OdinUser, error) {
	c := client.NewClient()

	// 发送请求
	endpoint := fmt.Sprintf("/user/%s", principalID)
	resp, err := c.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("获取用户信息失败: %w", err)
	}

	// 解析响应
	var user models.OdinUser
	if err := json.Unmarshal(resp, &user); err != nil {
		return nil, fmt.Errorf("解析用户信息失败: %w", err)
	}

	return &user, nil
}

// GetUserBalances 获取用户余额列表
// 获取指定用户拥有的所有代币余额
func GetUserBalances(principalID string) (*models.UserBalances, error) {
	c := client.NewClient()

	// 发送请求
	endpoint := fmt.Sprintf("/user/%s/balances", principalID)
	resp, err := c.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("获取用户余额失败: %w", err)
	}

	// 解析响应
	var balances models.UserBalances
	if err := json.Unmarshal(resp, &balances); err != nil {
		return nil, fmt.Errorf("解析用户余额失败: %w", err)
	}

	return &balances, nil
}

// GetUserTokenBalance 获取用户特定代币余额
// 获取指定用户特定代币的余额
func GetUserTokenBalance(principalID, tokenID string) (*models.UserBalance, error) {
	// 先获取所有余额
	balances, err := GetUserBalances(principalID)
	if err != nil {
		return nil, err
	}

	// 查找特定代币
	for _, balance := range balances.Data {
		if balance.ID == tokenID {
			return &balance, nil
		}
	}

	return nil, fmt.Errorf("未找到代币ID %s 的余额", tokenID)
}
