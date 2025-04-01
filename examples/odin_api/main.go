package main

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/MrHat365/odin-go/odin_api"
	"github.com/aviate-labs/agent-go/identity"
	"github.com/aviate-labs/agent-go/principal"
	"log"
	"time"
)

func PrettyPrint(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println(v)
		return
	}

	var out bytes.Buffer
	err = json.Indent(&out, b, "", "  ")
	if err != nil {
		fmt.Println(v)
		return
	}

	fmt.Println(out.String())
}

// IdentityAdapter 适配器结构体，用于桥接 Ed25519Identity 和 auth.Identity
type IdentityAdapter struct {
	*identity.Ed25519Identity
}

// GetPublicKey 实现 auth.Identity 接口
func (a *IdentityAdapter) GetPublicKey() []byte {
	return a.PublicKey()
}

// Sign 实现 auth.Identity 接口
func (a *IdentityAdapter) Sign(message []byte) ([]byte, error) {
	return a.Ed25519Identity.Sign(message), nil
}

// generateKeys 生成新的Ed25519密钥对
func generateKeys() (publicKey, privateKey []byte, err error) {
	// 使用crypto/ed25519包生成新的密钥对
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	return pubKey, privKey, nil
}

// loadKeys 从字符串加载密钥对
func loadKeys(pubKeyHex, privKeyHex string) (publicKey, privateKey []byte, err error) {
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

// saveKeys 将密钥对保存为十六进制字符串
func saveKeys(publicKey, privateKey []byte) (pubKeyHex, privKeyHex string) {
	return hex.EncodeToString(publicKey), hex.EncodeToString(privateKey)
}

// createPrincipalFromPublicKey 从公钥创建principal ID
func createPrincipalFromPublicKey(publicKey []byte) (string, error) {
	// 使用公钥创建principal
	p := principal.NewSelfAuthenticating(publicKey)
	return p.String(), nil
}

func main() {
	// 创建新的 Odin API 客户端
	client := odin_api.NewClient()

	pubKeyHex := "5d67c2887b8262c91af231d1a0b903efaebeb6e77dc45b8c898180228bb09318"
	privKeyHex := "3d4f2c109c3e2532a17119f9d136bcbf360744fba39601afbcee434b3525dd125d67c2887b8262c91af231d1a0b903efaebeb6e77dc45b8c898180228bb09318"
	loadedPubKey, loadedPrivKey, err := loadKeys(pubKeyHex, privKeyHex)
	if err != nil {
		log.Fatalf("加载密钥失败: %v", err)
	}
	log.Println(loadedPubKey)
	log.Println(loadedPrivKey)
	id, err := identity.NewEd25519Identity(loadedPubKey, loadedPrivKey)
	if err != nil {
		log.Fatalf("创建身份失败: %v", err)
	}

	// 演示 1: 创建身份并进行身份验证
	identityAdapter := IdentityAdapter{id}
	toekn := authenticateDemo(identityAdapter)

	client.SetToken(toekn)

	// 演示 2: 获取比特币价格
	//getBitcoinPriceDemo(client)
	//
	//// 演示 3: 获取市值最高的代币
	//topTokens := getTopTokensDemo(client)
	//
	//// 演示 4: 获取用户信息
	principalID, err := createPrincipalFromPublicKey(loadedPubKey)
	if err != nil {
		log.Fatalf("获取Principal ID失败: %v", err)
	}
	fmt.Printf("Principal ID: %s\n", principalID)
	getUserInfoDemo(client, principalID)

	//// 演示 5: 获取用户余额
	//getUserBalancesDemo(client, principalID)
	//
	//// 如果有市值最高的代币，继续演示获取特定代币的信息
	//tokenID := topTokens[0].ID
	//
	//// 演示 6: 获取特定代币的详细信息
	//getTokenDetailsDemo(client, tokenID)
	//
	//// 演示 7: 获取代币持有者信息
	//getTokenHoldersDemo(client, tokenID)
	//
	//// 演示 8: 使用授权令牌对代币发表评论
	//postCommentDemo(client, principalID, tokenID)
	//
	//// 演示 9: 更改用户名（需要授权令牌）
	//changeUsernameDemo(client, principalID, "NewUsername"+fmt.Sprintf("%d", time.Now().Unix()))
}

// authenticateDemo 演示如何创建身份并进行身份验证
// 返回创建的身份、授权令牌和用户ID
func authenticateDemo(id IdentityAdapter) string {
	token, err := odin_api.AuthIdentity(&id)
	if err != nil {
		log.Fatalf("%v", err)
	}

	// 从返回的令牌中提取principalID
	// 注意：这里假设令牌格式为 "Bearer {principalID}.{其他部分}"
	// 实际应用中可能需要调整提取方法
	fmt.Println("3. 成功获取授权令牌")

	fmt.Printf("   授权令牌: %s\n", token)

	return token
}

// getBitcoinPriceDemo 演示如何获取比特币价格
func getBitcoinPriceDemo(client *odin_api.Client) {
	fmt.Println("\n=== 演示 2: 获取比特币价格 ===")
	fmt.Println("调用API获取当前比特币价格...")

	btcInfo, err := client.GetBTCPrice()
	if err != nil {
		log.Printf("获取比特币价格失败: %v", err)
		return
	}

	fmt.Printf("比特币价格: $%.2f\n", btcInfo.Amount)
	fmt.Printf("时间: %s\n", btcInfo.Datetime.Format(time.RFC3339))
}

// getTopTokensDemo 演示如何获取市值最高的代币
func getTopTokensDemo(client *odin_api.Client) []odin_api.TokenDetail {
	fmt.Println("\n=== 演示 3: 获取市值最高的代币 ===")
	fmt.Println("调用API获取市值最高的代币列表...")

	tokens, err := client.GetTokensByHighestMarketcap()
	if err != nil {
		log.Printf("获取代币列表失败: %v", err)
		return nil
	}

	fmt.Printf("获取到 %d 个代币\n", len(tokens.Data))

	// 显示前5个代币的基本信息
	limit := 5
	if len(tokens.Data) < limit {
		limit = len(tokens.Data)
	}

	for i := 0; i < limit; i++ {
		token := tokens.Data[i]
		// 计算实际价格（考虑代币的小数位数）
		price := float64(token.Price) / float64(10^token.Decimals)
		marketcap := float64(token.Marketcap) / float64(10^token.Decimals)

		fmt.Printf("%d. %s (%s)\n", i+1, token.Name, token.Ticker)
		fmt.Printf("   价格: $%.6f\n", price)
		fmt.Printf("   市值: $%.2f\n", marketcap)
		fmt.Printf("   持有者数量: %d\n", token.HolderCount)
		fmt.Printf("   ID: %s\n", token.ID)
		fmt.Println()
	}

	return tokens.Data
}

// getUserInfoDemo 演示如何获取用户信息
func getUserInfoDemo(client *odin_api.Client, principalID string) {
	fmt.Println("\n=== 演示 4: 获取用户信息 ===")
	fmt.Printf("调用API获取用户ID %s 的信息...\n", principalID)

	userInfo, err := client.GetOdinFunUser(principalID)
	if err != nil {
		log.Printf("获取用户信息失败: %v", err)
		return
	}
	PrettyPrint(userInfo)
}

// getUserBalancesDemo 演示如何获取用户余额
func getUserBalancesDemo(client *odin_api.Client, principalID string) {
	fmt.Println("\n=== 演示 5: 获取用户余额 ===")
	fmt.Printf("调用API获取用户ID %s 的余额信息...\n", principalID)

	balances, err := client.GetUserBalances(principalID)
	if err != nil {
		log.Printf("获取用户余额失败: %v", err)
		return
	}

	fmt.Printf("获取到 %d 个代币余额\n", len(balances.Data))

	// 显示用户持有的代币余额
	for i, balance := range balances.Data {
		// 计算实际余额（考虑代币的小数位数）
		actualBalance := float64(balance.Balance) / float64(10^balance.Decimals)

		fmt.Printf("%d. %s (%s)\n", i+1, balance.Name, balance.Ticker)
		fmt.Printf("   余额: %.8f\n", actualBalance)
		fmt.Printf("   代币ID: %s\n", balance.ID)
		fmt.Println()
	}
}

// getTokenDetailsDemo 演示如何获取特定代币的详细信息
func getTokenDetailsDemo(client *odin_api.Client, tokenID string) {
	fmt.Println("\n=== 演示 6: 获取特定代币详情 ===")
	fmt.Printf("调用API获取代币ID %s 的详细信息...\n", tokenID)

	token, err := client.GetOdinFunToken(tokenID)
	if err != nil {
		log.Printf("获取代币详情失败: %v", err)
		return
	}

	// 计算实际价格（考虑代币的小数位数）
	price := float64(token.Price) / float64(10^token.Decimals)
	marketcap := float64(token.Marketcap) / float64(10^token.Decimals)

	fmt.Printf("名称: %s (%s)\n", token.Name, token.Ticker)
	fmt.Printf("描述: %s\n", token.Description)
	fmt.Printf("价格: $%.6f\n", price)
	fmt.Printf("市值: $%.2f\n", marketcap)
	fmt.Printf("总供应量: %d\n", token.TotalSupply)
	fmt.Printf("持有者数量: %d\n", token.HolderCount)
	fmt.Printf("创建者: %s\n", token.Creator)
	fmt.Printf("创建时间: %s\n", token.CreatedTime.Format(time.RFC3339))

	if token.Website != "" {
		fmt.Printf("网站: %s\n", token.Website)
	}
	if token.Twitter != "" {
		fmt.Printf("Twitter: %s\n", token.Twitter)
	}
	if token.Telegram != "" {
		fmt.Printf("Telegram: %s\n", token.Telegram)
	}
}

// getTokenHoldersDemo 演示如何获取代币持有者信息
func getTokenHoldersDemo(client *odin_api.Client, tokenID string) {
	fmt.Println("\n=== 演示 7: 获取代币持有者信息 ===")
	fmt.Printf("调用API获取代币ID %s 的持有者信息...\n", tokenID)

	holders, err := client.GetHolders(tokenID)
	if err != nil {
		log.Printf("获取持有者信息失败: %v", err)
		return
	}

	fmt.Printf("获取到 %d 个持有者\n", len(holders.Data))

	// 显示持有者信息
	for i, holder := range holders.Data {
		fmt.Printf("%d. 用户名: %s\n", i+1, holder.UserUsername)
		fmt.Printf("   用户ID: %s\n", holder.User)
		fmt.Printf("   持有量: %d\n", holder.Balance)
		fmt.Println()
	}
}

// postCommentDemo 演示如何对代币发表评论
func postCommentDemo(client *odin_api.Client, principalID, tokenID string) {
	fmt.Println("\n=== 演示 8: 发表代币评论 ===")
	fmt.Printf("对代币ID %s 发表评论...\n", tokenID)

	comment := "这是一个通过API发表的测试评论！" + fmt.Sprintf(" (时间戳: %d)", time.Now().Unix())

	result, err := client.PostComment(comment, principalID, tokenID)
	if err != nil {
		log.Printf("发表评论失败: %v", err)
		return
	}

	fmt.Printf("评论发表结果: %s\n", result)
}

// changeUsernameDemo 演示如何更改用户名
func changeUsernameDemo(client *odin_api.Client, principalID, newUsername string) {
	fmt.Println("\n=== 演示 9: 更改用户名 ===")
	fmt.Printf("将用户ID %s 的用户名更改为 %s...\n", principalID, newUsername)

	user, err := client.ChangeUsername(newUsername, principalID, client.Token)
	if err != nil {
		log.Printf("更改用户名失败: %v", err)
		return
	}

	fmt.Printf("用户名更改成功！新用户名: %s\n", user.Username)
}
