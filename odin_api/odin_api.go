// Package odin_api 提供与Odin.fun平台交互的Go语言API接口
//
// 该包实现了所有OdinFunAPI.cs中的功能，包括身份验证、用户操作、代币操作和市场数据获取
package odin_api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
)

// AuthRequest 身份验证请求结构
type AuthRequest struct {
	PublicKey string `json:"publickey"`
	Timestamp string `json:"timestamp"`
	Signature string `json:"signature"`
	Referrer  string `json:"referrer"`
}

// Identity 身份验证和身份注册请求
// identity参数需要实现GetPublicKey()和Sign()方法
type Identity interface {
	GetPublicKey() []byte
	Sign(message []byte) ([]byte, error)
}

// 身份验证相关功能

// AuthIdentity 身份验证和身份注册请求
func (c *Client) AuthIdentity(identity Identity) (string, error) {
	// 获取当前时间戳
	now := time.Now().UnixMilli()
	timestamp := fmt.Sprintf("%d", now)

	// 签名时间戳
	signature, err := identity.Sign([]byte(timestamp))
	if err != nil {
		return "", fmt.Errorf("签名生成失败: %w", err)
	}

	// 创建授权请求
	authReq := AuthRequest{
		PublicKey: base64.StdEncoding.EncodeToString(identity.GetPublicKey()),
		Timestamp: timestamp,
		Signature: base64.StdEncoding.EncodeToString(signature),
		Referrer:  "zg8khi8rz0",
	}

	// 发送请求
	resp, err := c.Post("/auth", authReq)
	if err != nil {
		return "", fmt.Errorf("身份验证请求失败: %w", err)
	}
	// 解析响应
	var authToken struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal(resp, &authToken); err != nil {
		return "", fmt.Errorf("解析授权令牌失败: %w", err)
	}

	return authToken.Token, nil
}

// 用户相关功能

// ChangeUsername 使用授权令牌更改用户名
func (c *Client) ChangeUsername(username, principalID, authToken string) (*OdinUser, error) {
	// 创建表单数据
	formData := map[string]string{
		"username": username,
	}

	// 发送请求
	endpoint := fmt.Sprintf("/user/profile?user=%s", principalID)
	resp, err := c.PostMultipart(endpoint, formData)
	if err != nil {
		return nil, fmt.Errorf("更改用户名请求失败: %w", err)
	}

	var odinUser *OdinUser

	if err := json.Unmarshal(resp, &odinUser); err != nil {
		return nil, fmt.Errorf("解析失败: %w", err)
	}

	return odinUser, nil
}

// GetOdinFunUser 获取Odin.fun用户信息
func (c *Client) GetOdinFunUser(principalID string) (*OdinUser, error) {
	// 发送请求
	endpoint := fmt.Sprintf("/odinUser/%s", principalID)
	resp, err := c.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("获取用户信息失败: %w", err)
	}

	// 解析响应
	var odinUser *OdinUser

	if err := json.Unmarshal(resp, &odinUser); err != nil {
		return nil, fmt.Errorf("解析用户信息失败: %w", err)
	}
	fmt.Println(odinUser)
	return odinUser, nil
}

// GetUserBalances 获取用户余额列表
func (c *Client) GetUserBalances(principalID string) (*OdinUserBalance, error) {
	// 发送请求
	endpoint := fmt.Sprintf("/user/%s/balances", principalID)
	resp, err := c.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("获取用户余额失败: %w", err)
	}

	// 解析响应
	var balances OdinUserBalance
	if err := json.Unmarshal(resp, &balances); err != nil {
		return nil, fmt.Errorf("解析用户余额失败: %w", err)
	}

	return &balances, nil
}

// GetUserTokenBalance 获取用户特定代币余额
func (c *Client) GetUserTokenBalance(principalID, tokenID string) (*BalanceDetail, error) {
	balances, err := c.GetUserBalances(principalID)
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

// 代币相关功能

// CommentRequest 发表评论请求结构
type CommentRequest struct {
	Message string `json:"message"`
}

// PostComment 发表评论
func (c *Client) PostComment(commentMessage, principalID, tokenID string) (string, error) {
	// 创建评论请求
	commentReq := CommentRequest{
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
func (c *Client) GetOdinFunTokens() (*OdinFunTokens, error) {
	// 发送请求
	endpoint := "/tokens?sort=last_action_time%3Adesc&page=1&limit=100"
	resp, err := c.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("获取代币列表失败: %w", err)
	}

	// 解析响应
	var tokens OdinFunTokens
	if err := json.Unmarshal(resp, &tokens); err != nil {
		return nil, fmt.Errorf("解析代币列表失败: %w", err)
	}

	return &tokens, nil
}

// GetTokensByHighestMarketcap 获取市值最高的Odin.fun代币
func (c *Client) GetTokensByHighestMarketcap() (*OdinFunTokens, error) {
	// 发送请求
	endpoint := "/tokens?sort=marketcap%3Adesc&page=1&limit=25"
	resp, err := c.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("获取代币列表失败: %w", err)
	}

	// 解析响应
	var tokens OdinFunTokens
	if err := json.Unmarshal(resp, &tokens); err != nil {
		return nil, fmt.Errorf("解析代币列表失败: %w", err)
	}

	return &tokens, nil
}

// GetHolders 获取代币持有者
func (c *Client) GetHolders(id string) (*Holders, error) {
	// 发送请求
	endpoint := fmt.Sprintf("/token/%s/owners?page=1&limit=10", id)
	resp, err := c.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("获取持有者列表失败: %w", err)
	}

	// 解析响应
	var holders Holders
	if err := json.Unmarshal(resp, &holders); err != nil {
		return nil, fmt.Errorf("解析持有者列表失败: %w", err)
	}

	return &holders, nil
}

// GetOdinFunToken 获取特定的Odin.fun代币
func (c *Client) GetOdinFunToken(id string) (*TokenDetail, error) {
	// 发送请求
	endpoint := fmt.Sprintf("/token/%s", id)
	resp, err := c.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("获取代币信息失败: %w", err)
	}

	// 解析响应
	var tokenResponse TokenDetail

	if err := json.Unmarshal(resp, &tokenResponse); err != nil {
		return nil, fmt.Errorf("解析代币信息失败: %w", err)
	}

	return &tokenResponse, nil
}

type TokenTarget struct {
	Id                  string `json:"id"`
	LastActionTimestamp int64  `json:"last_action_timestamp"`
}

// GetOdinFunTrades 获取特定代币的交易历史
func (c *Client) GetOdinFunTrades(target TokenTarget) (*TokenTraders, error) {
	// 发送请求
	endpoint := fmt.Sprintf("/token/%s/trades?page=1&limit=9999&time_min=%d", target.Id, target.LastActionTimestamp)
	resp, err := c.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("获取代币交易历史失败: %w", err)
	}

	// 解析响应
	var trades TokenTraders
	if err := json.Unmarshal(resp, &trades); err != nil {
		return nil, fmt.Errorf("解析代币交易历史失败: %w", err)
	}

	return &trades, nil
}

// 服务相关功能

// GetBTCPrice 获取比特币当前价格信息
func (c *Client) GetBTCPrice() (*BTCInfo, error) {
	// 发送请求
	endpoint := "/currency/btc"
	resp, err := c.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("获取比特币价格失败: %w", err)
	}

	// 解析响应
	var btcInfo BTCInfo
	if err := json.Unmarshal(resp, &btcInfo); err != nil {
		return nil, fmt.Errorf("解析比特币价格信息失败: %w", err)
	}

	return &btcInfo, nil
}
