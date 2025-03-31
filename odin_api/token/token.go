package token

import (
	"encoding/json"
	"fmt"

	"github.com/MrHat365/odin-go/odin_api/client"
	"github.com/MrHat365/odin-go/odin_api/models"
)

// PostComment 发表评论
// 用户可以对指定代币发表评论
func PostComment(commentMessage, principalID, tokenID, authToken string) (string, error) {
	c := client.NewClient()
	c.SetToken(authToken)

	// 创建评论请求
	commentReq := models.CommentRequest{
		Message: commentMessage,
	}

	// 发送请求
	endpoint := fmt.Sprintf("/token/%s/comment?user=%s", tokenID, principalID)
	resp, err := c.Post(endpoint, commentReq)
	if err != nil {
		return "", fmt.Errorf("发表评论请求失败: %w", err)
	}

	return string(resp), nil
}

// GetOdinFunTokens 获取最近交易的Odin.fun代币
// 返回按最后活动时间排序的代币列表
func GetOdinFunTokens() (*models.OdinFunTokens, error) {
	c := client.NewClient()

	// 发送请求
	endpoint := "/tokens?sort=last_action_time%3Adesc&page=1&limit=100"
	resp, err := c.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("获取代币列表失败: %w", err)
	}

	// 解析响应
	var tokens models.OdinFunTokens
	if err := json.Unmarshal(resp, &tokens); err != nil {
		return nil, fmt.Errorf("解析代币列表失败: %w", err)
	}

	return &tokens, nil
}

// GetTokensByHighestMarketcap 获取市值最高的Odin.fun代币
// 返回按市值排序的代币列表
func GetTokensByHighestMarketcap() (*models.OdinFunTokens, error) {
	c := client.NewClient()

	// 发送请求
	endpoint := "/tokens?sort=marketcap%3Adesc&page=1&limit=25"
	resp, err := c.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("获取代币列表失败: %w", err)
	}

	// 解析响应
	var tokens models.OdinFunTokens
	if err := json.Unmarshal(resp, &tokens); err != nil {
		return nil, fmt.Errorf("解析代币列表失败: %w", err)
	}

	return &tokens, nil
}

// GetHolders 获取代币持有者
// 返回指定代币的持有者列表
func GetHolders(id string) (*models.Holders, error) {
	c := client.NewClient()

	// 发送请求
	endpoint := fmt.Sprintf("/token/%s/owners?page=1&limit=10", id)
	resp, err := c.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("获取持有者列表失败: %w", err)
	}

	// 解析响应
	var holders models.Holders
	if err := json.Unmarshal(resp, &holders); err != nil {
		return nil, fmt.Errorf("解析持有者列表失败: %w", err)
	}

	return &holders, nil
}

// GetOdinFunToken 获取特定的Odin.fun代币
// 返回指定ID的代币详细信息
func GetOdinFunToken(id string) (*models.TokenData, error) {
	c := client.NewClient()

	// 发送请求
	endpoint := fmt.Sprintf("/token/%s", id)
	resp, err := c.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("获取代币信息失败: %w", err)
	}

	// 解析响应
	var tokenResponse struct {
		Status  bool             `json:"status"`
		Code    int              `json:"code"`
		Message string           `json:"message"`
		Data    models.TokenData `json:"data"`
	}

	if err := json.Unmarshal(resp, &tokenResponse); err != nil {
		return nil, fmt.Errorf("解析代币信息失败: %w", err)
	}

	return &tokenResponse.Data, nil
}

// GetOdinFunTrades 获取特定代币的交易历史
// 返回从指定时间戳开始的代币交易列表
func GetOdinFunTrades(target models.TokenTarget) (*models.TokenTrades, error) {
	c := client.NewClient()

	// 发送请求
	endpoint := fmt.Sprintf("/token/%s/trades?page=1&limit=9999&time_min=%d", target.Id, target.LastActionTimestamp)
	resp, err := c.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("获取代币交易历史失败: %w", err)
	}

	// 解析响应
	var trades models.TokenTrades
	if err := json.Unmarshal(resp, &trades); err != nil {
		return nil, fmt.Errorf("解析代币交易历史失败: %w", err)
	}

	return &trades, nil
}
