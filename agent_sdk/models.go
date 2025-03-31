// Package agent_sdk 提供与Internet Computer上的AgentSdk智能合约交互的客户端
package agent_sdk

import (
	"encoding/json"
	"math/big"
)

// TokenID 表示代币的唯一标识符
type TokenID = string

// TokenAmount 表示代币数量，使用无界整数类型
type TokenAmount = *big.Int

// OptionalValue 表示一个可能为空的值
type OptionalValue[T any] struct {
	HasValue bool
	Value    T
}

// Token 表示一个代币的完整信息
type Token struct {
	ID          TokenID     `json:"id"`
	Name        string      `json:"name"`
	Symbol      string      `json:"symbol"`
	TotalSupply TokenAmount `json:"totalSupply"`
	Decimals    uint8       `json:"decimals"`
	Owner       string      `json:"owner"`
	// 其他属性可根据需要添加
}

// LockedTokenState 表示锁定的代币状态
type LockedTokenState struct {
	Amount    TokenAmount `json:"amount"`
	UnlockAt  uint64      `json:"unlockAt"`
	TokenID   TokenID     `json:"tokenId"`
	LockOwner string      `json:"lockOwner"`
}

// Operation 表示一个操作记录
type Operation struct {
	Timestamp uint64              `json:"timestamp"`
	Caller    string              `json:"caller"`
	Op        string              `json:"op"`
	Details   map[string]TokenIDs `json:"details"`
}

// TokenIDs 表示代币ID的列表
type TokenIDs []TokenID

// MarshalJSON 实现自定义JSON序列化
func (t TokenIDs) MarshalJSON() ([]byte, error) {
	return json.Marshal([]string(t))
}

// OperationAndId 表示操作及其ID
type OperationAndId struct {
	Op Operation   `json:"op"`
	ID TokenAmount `json:"id"`
}

// AddRequest 表示添加代币的请求
type AddRequest struct {
	TokenID TokenID     `json:"tokenId"`
	Reserve TokenAmount `json:"reserve"`
	Fee     TokenAmount `json:"fee"`
}

// AddResponse 表示添加代币的响应
type AddResponse struct {
	OK  *TokenAmount `json:"ok,omitempty"`
	Err *string      `json:"err,omitempty"`
}

// EtchRequest 表示刻印代币的请求
type EtchRequest struct {
	TokenID     TokenID     `json:"tokenId"`
	Name        string      `json:"name"`
	Symbol      string      `json:"symbol"`
	TotalSupply TokenAmount `json:"totalSupply"`
	Decimals    uint8       `json:"decimals"`
}

// EtchResponse 表示刻印代币的响应
type EtchResponse struct {
	OK  *TokenAmount `json:"ok,omitempty"`
	Err *string      `json:"err,omitempty"`
}

// LiquidityRequest 表示流动性操作的请求
type LiquidityRequest struct {
	TokenID      TokenID     `json:"tokenId"`
	Amount       TokenAmount `json:"amount"`
	Operation    string      `json:"operation"` // "add" 或 "remove"
	MinimumPrice TokenAmount `json:"minimumPrice,omitempty"`
}

// LiquidityResponse 表示流动性操作的响应
type LiquidityResponse struct {
	OK  *TokenAmount `json:"ok,omitempty"`
	Err *string      `json:"err,omitempty"`
}

// MintRequest 表示铸造代币的请求
type MintRequest struct {
	TokenID TokenID     `json:"tokenId"`
	To      string      `json:"to"`
	Amount  TokenAmount `json:"amount"`
}

// MintResponse 表示铸造代币的响应
type MintResponse struct {
	OK  *TokenAmount `json:"ok,omitempty"`
	Err *string      `json:"err,omitempty"`
}

// TradeRequest 表示交易代币的请求
type TradeRequest struct {
	TokenID        TokenID     `json:"tokenId"`
	Amount         TokenAmount `json:"amount"`
	Operation      string      `json:"operation"` // "buy" 或 "sell"
	MaxSlippage    TokenAmount `json:"maxSlippage,omitempty"`
	ExpectedAmount TokenAmount `json:"expectedAmount,omitempty"`
}

// TradeResponse 表示交易代币的响应
type TradeResponse struct {
	OK  *TokenAmount `json:"ok,omitempty"`
	Err *string      `json:"err,omitempty"`
}

// WithdrawRequest 表示提取代币的请求
type WithdrawRequest struct {
	TokenID TokenID     `json:"tokenId"`
	Amount  TokenAmount `json:"amount"`
	To      string      `json:"to,omitempty"`
}

// WithdrawResponse 表示提取代币的响应
type WithdrawResponse struct {
	OK  *TokenAmount `json:"ok,omitempty"`
	Err *string      `json:"err,omitempty"`
}
