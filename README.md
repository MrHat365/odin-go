# Odin-Go：Odin.fun 与 Internet Computer 客户端

Odin-Go 是一个用 Go 语言编写的客户端库集合，用于与 Odin.fun 平台和 Internet Computer 上的 AgentSdk 智能合约进行交互。该项目包含两个主要组件：

1. `agent_sdk`：与 Internet Computer 上的 AgentSdk 智能合约交互的客户端库
2. `odin_api`：与 Odin.fun 平台 API 交互的客户端库

## 功能特性

### agent_sdk

- 查询代币余额和信息
- 获取锁定的代币信息和操作记录
- 添加和管理代币流动性
- 交易和转移代币
- 铸造和提取代币
- 代币存款和提款
- 辅助计算和转换函数

### odin_api

- 身份验证和用户注册
- 用户管理（更改用户名、获取用户信息）
- 代币查询（获取代币列表、代币详情、持有者信息）
- 交易查询（获取交易历史）
- 市场数据（获取比特币价格）

## 安装

```bash
go get github.com/MrHat365/odin-go
```

## 快速开始

### 使用 agent_sdk

```go
package main

import (
	"fmt"
	"log"

	"github.com/aviate-labs/agent-go"
	"github.com/MrHat365/odin-go/agent_sdk"
)

func main() {
	// 创建一个匿名身份的 Agent
	ag, err := agent.New(agent.DefaultConfig)
	if err != nil {
		log.Fatalf("创建 Agent 失败: %v", err)
	}
	
	// 创建 AgentSdk 客户端，使用默认的 Canister ID
	client, err := agent_sdk.New(ag, "")
	if err != nil {
		log.Fatalf("创建 AgentSdk 客户端失败: %v", err)
	}
	
	// 查询代币余额
	balance, err := client.GetBalance("myAccount", "principal", "myTokenID")
	if err != nil {
		log.Printf("查询余额失败: %v", err)
	} else {
		fmt.Printf("余额: %s\n", (*balance).String())
	}
}
```

### 使用 odin_api

```go
package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
	"log"

	"github.com/MrHat365/odin-go/odin_api"
)

// 实现 auth.Identity 接口的结构
type Ed25519Identity struct {
	PrivateKey ed25519.PrivateKey
	PublicKey  ed25519.PublicKey
}

func (i *Ed25519Identity) GetPublicKey() []byte {
	return i.PublicKey
}

func (i *Ed25519Identity) Sign(message []byte) ([]byte, error) {
	signature := ed25519.Sign(i.PrivateKey, message)
	return signature, nil
}

func main() {
	// 创建新身份
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatalf("生成密钥对失败: %v", err)
	}
	
	identity := &Ed25519Identity{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}

	// 身份验证
	authToken, err := odin_api.AuthIdentity(identity)
	if err != nil {
		log.Fatalf("身份验证失败: %v", err)
	}

	// 保存凭据
	token := authToken.Data.Token
	principalID := authToken.Data.PrincipalID
	
	// 获取比特币价格
	btcInfo, err := odin_api.GetBTCPrice()
	if err != nil {
		log.Fatalf("获取比特币价格失败: %v", err)
	}
	fmt.Printf("比特币当前价格: $%.2f\n", btcInfo.Data.Price)
	
	// 获取用户信息
	userInfo, err := odin_api.GetOdinFunUser(principalID)
	if err != nil {
		log.Fatalf("获取用户信息失败: %v", err)
	}
	fmt.Printf("用户名: %s\n", userInfo.Data.Username)
}
```

## API 文档

### agent_sdk

#### 创建客户端

```go
// 创建新的 AgentSdk 客户端
client, err := agent_sdk.New(agent, canisterID)
```

如果未提供 canisterID，将使用默认值 "z2vm5-gaaaa-aaaaj-azw6q-cai"。

#### 查询方法

```go
// 获取代币余额
balance, err := client.GetBalance(accountId, accountType, tokenId)

// 获取锁定的代币信息
lockedTokens, err := client.GetLockedTokens(accountId)

// 获取特定操作的详细信息
operation, err := client.GetOperation(accountId, operationId)

// 获取一系列操作记录
operations, err := client.GetOperations(startId, endId)

// 获取统计信息
stats, err := client.GetStats(statsType)

// 获取代币信息
token, err := client.GetToken(accountId, tokenId)

// 获取代币索引
index, err := client.GetTokenIndex(tokenId)
```

#### 更新方法

```go
// 添加代币
response, err := client.TokenAdd(addRequest)

// 存入代币
result, err := client.TokenDeposit(tokenId, amount)

// 铸造代币
response, err := client.TokenEtch(etchRequest)

// 处理代币流动性
response, err := client.TokenLiquidity(liquidityRequest)

// 铸造代币到指定地址
response, err := client.TokenMint(mintRequest)

// 交易代币
response, err := client.TokenTrade(tradeRequest)

// 提取代币
response, err := client.TokenWithdraw(withdrawRequest)
```

### odin_api

#### 身份验证

```go
// 身份验证和注册
authToken, err := odin_api.AuthIdentity(identity)
```

#### 用户相关

```go
// 更改用户名
result, err := odin_api.ChangeUsername(username, principalID, token)

// 获取用户信息
userInfo, err := odin_api.GetOdinFunUser(principalID)

// 获取用户余额列表
balances, err := odin_api.GetUserBalances(principalID)

// 获取用户特定代币余额
balance, err := odin_api.GetUserTokenBalance(principalID, tokenID)
```

#### 代币相关

```go
// 发表评论
result, err := odin_api.PostComment(commentMessage, principalID, tokenID, authToken)

// 获取最近交易的代币
tokens, err := odin_api.GetOdinFunTokens()

// 获取市值最高的代币
tokens, err := odin_api.GetTokensByHighestMarketcap()

// 获取代币持有者
holders, err := odin_api.GetHolders(tokenID)

// 获取特定代币信息
token, err := odin_api.GetOdinFunToken(tokenID)

// 获取代币交易历史
trades, err := odin_api.GetOdinFunTrades(tokenTarget)
```

#### 其他服务

```go
// 获取比特币价格
btcInfo, err := odin_api.GetBTCPrice()
```

## 密钥和身份管理

在 Internet Computer 上，身份由密钥对表示，Principal ID 是用户的唯一标识符。以下是管理密钥和身份的示例代码：

```go
// 生成新的 Ed25519 密钥对
pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)

// 创建 Ed25519 身份
id, err := identity.NewEd25519Identity(publicKey, privateKey)

// 使用特定身份创建 Agent
config := agent.Config{
	Identity: id,
}
ag, err := agent.New(config)
```

## 工具函数

`agent_sdk` 包提供了一系列辅助工具函数：

```go
// 将 satoshis 转换为 BTC
btc := agent_sdk.ConvertToBTC(satoshisValue)

// 将 satoshis 转换为代币数量
amount := agent_sdk.ConvertToTokenAmount(satoshisValue)

// 计算两个数值之间的百分比差异
diff := agent_sdk.CalculatePercentDifference(a, b)

// 创建新的 TokenAmount
amount := agent_sdk.NewTokenAmount(value)

// 从字符串创建 TokenAmount
amount, success := agent_sdk.NewTokenAmountFromString(valueStr)

// 将 TokenAmount 转换为 float64
float := agent_sdk.TokenAmountToFloat64(amount)
```

## 安全注意事项

使用本库时需要注意以下安全风险：

1. **密钥管理**：私钥用于签名交易，必须安全存储。私钥泄露将导致资产被盗。
   - 建议使用硬件钱包或安全的密钥管理系统
   - 不要在代码或配置文件中硬编码私钥
   - 考虑使用环境变量或加密存储解决方案

2. **交易验证**：在执行转账或交易前，务必验证交易参数。
   - 验证接收地址的正确性
   - 确认转账金额无误
   - 实现多重签名或交易确认机制

3. **API 请求安全**：
   - 所有请求都应使用 HTTPS
   - 验证 API 响应的完整性
   - 实现速率限制以防止 API 滥用

4. **错误处理**：
   - 妥善处理所有错误情况
   - 不要在生产环境中暴露详细错误信息
   - 实现故障恢复机制

5. **依赖项风险**：
   - 定期更新依赖项以修复已知漏洞
   - 审查第三方库的安全历史
   - 考虑使用锁定版本的依赖项
