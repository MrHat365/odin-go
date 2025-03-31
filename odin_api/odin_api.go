// Package odin_api 提供与Odin.fun平台交互的Go语言API接口
//
// 该包实现了所有OdinFunAPI.cs中的功能，包括身份验证、用户操作、代币操作和市场数据获取
package odin_api

import (
	"github.com/MrHat365/odin-go/odin_api/auth"
	"github.com/MrHat365/odin-go/odin_api/models"
	"github.com/MrHat365/odin-go/odin_api/services"
	"github.com/MrHat365/odin-go/odin_api/token"
	"github.com/MrHat365/odin-go/odin_api/user"
)

// 身份验证相关功能

// AuthIdentity 身份验证和身份注册请求
func AuthIdentity(identity auth.Identity) (*models.AuthToken, error) {
	return auth.AuthIdentity(identity)
}

// 用户相关功能

// ChangeUsername 使用授权令牌更改用户名
func ChangeUsername(username, principalID, authToken string) (string, error) {
	return user.ChangeUsername(username, principalID, authToken)
}

// GetOdinFunUser 获取Odin.fun用户信息
func GetOdinFunUser(principalID string) (*models.OdinUser, error) {
	return user.GetOdinFunUser(principalID)
}

// GetUserBalances 获取用户余额列表
func GetUserBalances(principalID string) (*models.UserBalances, error) {
	return user.GetUserBalances(principalID)
}

// GetUserTokenBalance 获取用户特定代币余额
func GetUserTokenBalance(principalID, tokenID string) (*models.UserBalance, error) {
	return user.GetUserTokenBalance(principalID, tokenID)
}

// 代币相关功能

// PostComment 发表评论
func PostComment(commentMessage, principalID, tokenID, authToken string) (string, error) {
	return token.PostComment(commentMessage, principalID, tokenID, authToken)
}

// GetOdinFunTokens 获取最近交易的Odin.fun代币
func GetOdinFunTokens() (*models.OdinFunTokens, error) {
	return token.GetOdinFunTokens()
}

// GetTokensByHighestMarketcap 获取市值最高的Odin.fun代币
func GetTokensByHighestMarketcap() (*models.OdinFunTokens, error) {
	return token.GetTokensByHighestMarketcap()
}

// GetHolders 获取代币持有者
func GetHolders(id string) (*models.Holders, error) {
	return token.GetHolders(id)
}

// GetOdinFunToken 获取特定的Odin.fun代币
func GetOdinFunToken(id string) (*models.TokenData, error) {
	return token.GetOdinFunToken(id)
}

// GetOdinFunTrades 获取特定代币的交易历史
func GetOdinFunTrades(target models.TokenTarget) (*models.TokenTrades, error) {
	return token.GetOdinFunTrades(target)
}

// 服务相关功能

// GetBTCPrice 获取比特币当前价格信息
func GetBTCPrice() (*models.BTCInfo, error) {
	return services.GetBTCPrice()
}
