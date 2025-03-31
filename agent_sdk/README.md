# AgentSdk客户端库

这是一个用Go语言编写的客户端库，用于与Internet Computer上的AgentSdk智能合约进行交互。该库基于[agent-go](https://github.com/aviate-labs/agent-go)开发，提供了一套完整的API来操作AgentSdk平台上的代币。

## 功能特性

- 查询代币余额和信息
- 获取操作记录和统计数据
- 添加和管理代币流动性
- 交易和转移代币
- 铸造和提取代币
- 代币存款和提款
- 辅助计算和转换函数

## 安装

```bash
go get github.com/wyfz/odin/AgentSdk
```

## 快速开始

以下是一个简单的示例，演示如何使用该库与AgentSdk智能合约交互：

```go
package main

import (
	"fmt"
	"log"

	"github.com/aviate-labs/agent-go"
	"github.com/wyfz/odin/AgentSdk"
)

func main() {
	// 创建一个匿名身份的Agent
	ag, err := agent.New(agent.DefaultConfig)
	if err != nil {
		log.Fatalf("创建Agent失败: %v", err)
	}
	
	// 创建AgentSdk客户端，使用默认的Canister ID
	client, err := AgentSdk.New(ag, "")
	if err != nil {
		log.Fatalf("创建AgentSdk客户端失败: %v", err)
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

## 密钥和身份管理

在Internet Computer上，身份由密钥对表示，Principal ID是用户的唯一标识符。以下是创建和管理密钥对和Principal ID的示例：

### 生成新的Ed25519密钥对

```go
import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
)

// 生成新的密钥对
func generateKeys() (publicKey, privateKey []byte, err error) {
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	return pubKey, privKey, nil
}

// 将密钥保存为十六进制字符串
func saveKeys(publicKey, privateKey []byte) (pubKeyHex, privKeyHex string) {
	return hex.EncodeToString(publicKey), hex.EncodeToString(privateKey)
}

// 使用示例
pubKey, privKey, _ := generateKeys()
pubKeyHex, privKeyHex := saveKeys(pubKey, privKey)
fmt.Printf("公钥: %s\n", pubKeyHex)
fmt.Printf("私钥: %s\n", privKeyHex)
```

### 从十六进制字符串加载密钥对

```go
// 从十六进制字符串加载密钥对
func loadKeys(pubKeyHex, privKeyHex string) ([]byte, []byte, error) {
	pubKey, err := hex.DecodeString(pubKeyHex)
	if err != nil {
		return nil, nil, fmt.Errorf("解析公钥失败: %w", err)
	}

	privKey, err := hex.DecodeString(privKeyHex)
	if err != nil {
		return nil, nil, fmt.Errorf("解析私钥失败: %w", err)
	}

	return pubKey, privKey, nil
}
```

### 从公钥创建Principal ID

```go
import (
	"github.com/aviate-labs/agent-go/principal"
)

// 从公钥创建principal ID
func createPrincipalFromPublicKey(publicKey []byte) (string, error) {
	p := principal.NewSelfAuthenticating(publicKey)
	return p.String(), nil
}
```

### 保存和加载密钥对

```go
import (
	"os"
	"strings"
)

// 保存密钥对到文件
func saveKeysToFile(filename string, pubKeyHex, privKeyHex string) error {
	return os.WriteFile(filename, []byte(pubKeyHex+"\n"+privKeyHex), 0600)
}

// 从文件加载密钥对
func loadKeysFromFile(filename string) (pubKeyHex, privKeyHex string, err error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", "", err
	}
	
	lines := strings.Split(string(data), "\n")
	if len(lines) < 2 {
		return "", "", fmt.Errorf("文件格式不正确")
	}
	
	return lines[0], lines[1], nil
}
```

## 使用自定义身份

如果需要使用特定身份，可以按照以下方式创建Agent：

```go
import (
	"github.com/aviate-labs/agent-go"
	"github.com/aviate-labs/agent-go/identity"
)

// 创建Ed25519身份
id, err := identity.NewEd25519Identity(publicKey, privateKey)
if err != nil {
	log.Fatalf("创建身份失败: %v", err)
}

// 使用该身份创建Agent
config := agent.Config{
	Identity: id,
}
ag, err := agent.New(config)
```

## 连接到本地开发环境

如果需要连接到本地开发环境，可以使用以下配置：

```go
import (
	"net/url"
	"github.com/aviate-labs/agent-go"
)

u, _ := url.Parse("http://localhost:8000")
config := agent.Config{
	ClientConfig: &agent.ClientConfig{Host: u},
	FetchRootKey: true,
	DisableSignedQueryVerification: true,
}
ag, err := agent.New(config)
```

## API参考

### 创建客户端

```go
// 创建新的AgentSdk客户端
client, err := AgentSdk.New(agent, canisterID)
```

如果未提供canisterID，将使用默认值"z2vm5-gaaaa-aaaaj-azw6q-cai"。

### 查询方法

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

### 更新方法

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

### 辅助函数

```go
// 将satoshis转换为BTC
btc := AgentSdk.ConvertToBTC(satoshis)

// 将satoshis转换为代币数量
amount := AgentSdk.ConvertToTokenAmount(satoshis)

// 计算两个数值之间的百分比差异
diff := AgentSdk.CalculatePercentDifference(a, b)

// 从int64创建TokenAmount类型
amount := AgentSdk.NewTokenAmount(value)

// 从字符串创建TokenAmount类型
amount, success := AgentSdk.NewTokenAmountFromString(value)

// 将TokenAmount转换为float64类型
f := AgentSdk.TokenAmountToFloat64(amount)
```

## 完整示例

更多示例请查看[示例目录](./example)。

## 许可证

本项目采用MIT许可证。 