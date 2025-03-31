package agent_sdk

import (
	"errors"
	"fmt"

	"github.com/aviate-labs/agent-go"
	"github.com/aviate-labs/agent-go/principal"
)

// Client 是AgentSdk智能合约的Golang客户端
type Client struct {
	Agent      *agent.Agent
	CanisterID principal.Principal
}

// DefaultCanisterID 是AgentSdk智能合约的默认Canister ID
const DefaultCanisterID = "z2vm5-gaaaa-aaaaj-azw6q-cai"

// New 创建一个新的AgentSdk客户端
// 如果未提供canisterID，则使用默认值
func New(agent *agent.Agent, canisterID string) (*Client, error) {
	if agent == nil {
		return nil, errors.New("agent不能为空")
	}

	var canister principal.Principal
	var err error

	if canisterID == "" {
		canister, err = principal.Decode(DefaultCanisterID)
	} else {
		canister, err = principal.Decode(canisterID)
	}

	if err != nil {
		return nil, fmt.Errorf("无效的canister ID: %w", err)
	}

	return &Client{
		Agent:      agent,
		CanisterID: canister,
	}, nil
}

// GetBalance 获取账户的代币余额
// arg0: 账户标识符
// arg1: 账户类型
// arg2: 代币ID
func (c *Client) GetBalance(arg0, arg1 string, arg2 TokenID) (*TokenAmount, error) {
	// 组装参数
	args := []any{arg0, arg1, arg2}

	// 发送查询请求
	var response TokenAmount
	err := c.Agent.Query(c.CanisterID, "getBalance", args, []any{&response})
	if err != nil {
		return nil, fmt.Errorf("GetBalance请求失败: %w", err)
	}

	return &response, nil
}

// GetLockedTokens 获取锁定的代币信息
// arg0: 账户标识符
func (c *Client) GetLockedTokens(arg0 string) (*LockedTokenState, error) {
	// 组装参数
	args := []any{arg0}

	// 发送查询请求
	var response LockedTokenState
	err := c.Agent.Query(c.CanisterID, "getLockedTokens", args, []any{&response})
	if err != nil {
		return nil, fmt.Errorf("GetLockedTokens请求失败: %w", err)
	}

	return &response, nil
}

// GetOperation 获取特定操作的详细信息
// arg0: 账户标识符
// arg1: 操作ID
func (c *Client) GetOperation(arg0 string, arg1 TokenAmount) (*OptionalValue[Operation], error) {
	// 组装参数
	args := []any{arg0, arg1}

	// 发送查询请求
	var response OptionalValue[Operation]
	err := c.Agent.Query(c.CanisterID, "getOperation", args, []any{&response})
	if err != nil {
		return nil, fmt.Errorf("GetOperation请求失败: %w", err)
	}

	return &response, nil
}

// GetOperations 获取一系列操作记录
// arg0: 起始ID
// arg1: 结束ID
func (c *Client) GetOperations(arg0, arg1 TokenAmount) ([]OperationAndId, error) {
	// 组装参数
	args := []any{arg0, arg1}

	// 发送查询请求
	var response []OperationAndId
	err := c.Agent.Query(c.CanisterID, "getOperations", args, []any{&response})
	if err != nil {
		return nil, fmt.Errorf("GetOperations请求失败: %w", err)
	}

	return response, nil
}

// GetStats 获取统计信息
// arg0: 统计类型
func (c *Client) GetStats(arg0 string) (map[string]string, error) {
	// 组装参数
	args := []any{arg0}

	// 发送查询请求
	var response map[string]string
	err := c.Agent.Query(c.CanisterID, "getStats", args, []any{&response})
	if err != nil {
		return nil, fmt.Errorf("GetStats请求失败: %w", err)
	}

	return response, nil
}

// GetToken 获取代币信息
// arg0: 账户标识符
// tokenID: 代币ID
func (c *Client) GetToken(arg0 string, tokenID TokenID) (*OptionalValue[Token], error) {
	// 组装参数
	args := []any{arg0, tokenID}

	// 发送查询请求
	var response OptionalValue[Token]
	err := c.Agent.Query(c.CanisterID, "getToken", args, []any{&response})
	if err != nil {
		return nil, fmt.Errorf("GetToken请求失败: %w", err)
	}

	return &response, nil
}

// GetTokenIndex 获取代币索引
// tokenID: 代币ID
func (c *Client) GetTokenIndex(tokenID TokenID) (*TokenAmount, error) {
	// 组装参数
	args := []any{tokenID}

	// 发送查询请求
	var response TokenAmount
	err := c.Agent.Query(c.CanisterID, "getTokenIndex", args, []any{&response})
	if err != nil {
		return nil, fmt.Errorf("GetTokenIndex请求失败: %w", err)
	}

	return &response, nil
}

// TokenAdd 添加代币
// request: 添加代币请求
func (c *Client) TokenAdd(request AddRequest) (*AddResponse, error) {
	// 组装参数
	args := []any{request}

	// 发送更新请求
	var response AddResponse
	err := c.Agent.Call(c.CanisterID, "token_add", args, []any{&response})
	if err != nil {
		return nil, fmt.Errorf("TokenAdd请求失败: %w", err)
	}

	return &response, nil
}

// TokenDeposit 存入代币
// tokenID: 代币ID
// amount: 存入金额
func (c *Client) TokenDeposit(tokenID TokenID, amount TokenAmount) (*TokenAmount, error) {
	// 组装参数
	args := []any{tokenID, amount}

	// 发送更新请求
	var response TokenAmount
	err := c.Agent.Call(c.CanisterID, "token_deposit", args, []any{&response})
	if err != nil {
		return nil, fmt.Errorf("TokenDeposit请求失败: %w", err)
	}

	return &response, nil
}

// TokenEtch 铸造代币
// request: 铸造代币请求
func (c *Client) TokenEtch(request EtchRequest) (*EtchResponse, error) {
	// 组装参数
	args := []any{request}

	// 发送更新请求
	var response EtchResponse
	err := c.Agent.Call(c.CanisterID, "token_etch", args, []any{&response})
	if err != nil {
		return nil, fmt.Errorf("TokenEtch请求失败: %w", err)
	}

	return &response, nil
}

// TokenLiquidity 处理代币流动性
// request: 流动性请求
func (c *Client) TokenLiquidity(request LiquidityRequest) (*LiquidityResponse, error) {
	// 组装参数
	args := []any{request}

	// 发送更新请求
	var response LiquidityResponse
	err := c.Agent.Call(c.CanisterID, "token_liquidity", args, []any{&response})
	if err != nil {
		return nil, fmt.Errorf("TokenLiquidity请求失败: %w", err)
	}

	return &response, nil
}

// TokenMint 铸造代币到指定地址
// request: 铸造请求
func (c *Client) TokenMint(request MintRequest) (*MintResponse, error) {
	// 组装参数
	args := []any{request}

	// 发送更新请求
	var response MintResponse
	err := c.Agent.Call(c.CanisterID, "token_mint", args, []any{&response})
	if err != nil {
		return nil, fmt.Errorf("TokenMint请求失败: %w", err)
	}

	return &response, nil
}

// TokenTrade 交易代币
// request: 交易请求
func (c *Client) TokenTrade(request TradeRequest) (*TradeResponse, error) {
	// 组装参数
	args := []any{request}

	// 发送更新请求
	var response TradeResponse
	err := c.Agent.Call(c.CanisterID, "token_trade", args, []any{&response})
	if err != nil {
		return nil, fmt.Errorf("TokenTrade请求失败: %w", err)
	}

	return &response, nil
}

// TokenWithdraw 提取代币
// request: 提取请求
func (c *Client) TokenWithdraw(request WithdrawRequest) (*WithdrawResponse, error) {
	// 组装参数
	args := []any{request}

	// 发送更新请求
	var response WithdrawResponse
	err := c.Agent.Call(c.CanisterID, "token_withdraw", args, []any{&response})
	if err != nil {
		return nil, fmt.Errorf("TokenWithdraw请求失败: %w", err)
	}

	return &response, nil
}
