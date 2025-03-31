# Odin.fun API Go 客户端

这个包提供了与 Odin.fun 平台交互的 Go 语言 API 接口。它是从 C# 的 `OdinFunAPI.cs` 移植而来的完整实现。

## 功能特性

- 身份验证和注册
- 用户管理（更改用户名、获取用户信息）
- 代币查询（获取代币列表、代币详情、持有者信息）
- 交易查询（获取交易历史）
- 市场数据（获取比特币价格）

## 安装

```bash
go get github.com/wyfz/odin/odin_api
```

## 使用示例

### 初始化并进行身份验证

```go
import (
    "crypto/ed25519"
    "crypto/rand"
    "fmt"
    "github.com/wyfz/odin/odin_api"
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

// 创建新身份
publicKey, privateKey, _ := ed25519.GenerateKey(rand.Reader)
identity := &Ed25519Identity{
    PrivateKey: privateKey,
    PublicKey:  publicKey,
}

// 身份验证
authToken, err := odin_api.AuthIdentity(identity)
if err != nil {
    panic(err)
}

// 保存凭据
token := authToken.Data.Token
principalID := authToken.Data.PrincipalID
```

### 获取比特币价格

```go
btcInfo, err := odin_api.GetBTCPrice()
if err != nil {
    panic(err)
}
fmt.Printf("比特币当前价格: $%.2f\n", btcInfo.Data.Price)
```

### 获取市值最高的代币

```go
tokens, err := odin_api.GetTokensByHighestMarketcap()
if err != nil {
    panic(err)
}
for i, token := range tokens.Data {
    fmt.Printf("%s (%s) - 价格: $%.4f\n", token.Name, token.Symbol, token.Price)
}
```

### 更改用户名

```go
result, err := odin_api.ChangeUsername("NewUsername", principalID, token)
if err != nil {
    panic(err)
}
fmt.Println("更改用户名结果:", result)
```

### 获取用户信息

```go
userInfo, err := odin_api.GetOdinFunUser(principalID)
if err != nil {
    panic(err)
}
fmt.Printf("用户名: %s\n", userInfo.Data.Username)
```

## 包结构

- `models` - 数据模型和结构体定义
- `client` - HTTP 客户端实现
- `auth` - 身份验证相关功能
- `user` - 用户相关功能
- `token` - 代币相关功能
- `services` - 其他服务（如获取比特币价格）

## 注意事项

- 所有 API 调用都是同步的，但使用了 Go 的标准错误处理模式
- 身份验证需要实现 `auth.Identity` 接口的对象
- 某些 API 调用需要有效的授权令牌 