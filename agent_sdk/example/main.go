// 示例程序，展示如何使用AgentSdk客户端库
package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/aviate-labs/agent-go/identity"
	"github.com/aviate-labs/agent-go/principal"
)

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
	// 示例1: 创建新的密钥对
	fmt.Println("===== 示例1: 创建新的密钥对 =====")
	pubKey, privKey, err := generateKeys()
	if err != nil {
		log.Fatalf("生成密钥对失败: %v", err)
	}

	pubKeyHex, privKeyHex := saveKeys(pubKey, privKey)
	fmt.Printf("生成的公钥: %s\n", pubKeyHex)
	fmt.Printf("生成的私钥: %s\n", privKeyHex)

	// 示例2: 从十六进制字符串加载密钥
	fmt.Println("\n===== 示例2: 从十六进制字符串加载密钥 =====")
	loadedPubKey, loadedPrivKey, err := loadKeys(pubKeyHex, privKeyHex)
	if err != nil {
		log.Fatalf("加载密钥失败: %v", err)
	}
	fmt.Printf("加载的公钥长度: %d bytes\n", len(loadedPubKey))
	fmt.Printf("加载的私钥长度: %d bytes\n", len(loadedPrivKey))

	// 示例3: 创建身份并获取Principal ID
	fmt.Println("\n===== 示例3: 创建身份并获取Principal ID =====")
	id, err := identity.NewEd25519Identity(loadedPubKey, loadedPrivKey)
	if err != nil {
		log.Fatalf("创建身份失败: %v", err)
	}
	fmt.Println(id)

	// 直接从公钥获取Principal ID
	principalID, err := createPrincipalFromPublicKey(loadedPubKey)
	if err != nil {
		log.Fatalf("获取Principal ID失败: %v", err)
	}
	fmt.Printf("Principal ID: %s\n", principalID)

	// 示例4: 使用身份创建Agent并连接到AgentSdk合约
	//fmt.Println("\n===== 示例4: 使用身份创建Agent并连接到AgentSdk合约 =====")
	//config := agent.Config{
	//	Identity: id, // 使用刚创建的身份
	//}

	//ag, err := agent.New(config)
	//if err != nil {
	//	log.Fatalf("创建Agent失败: %v", err)
	//}
	//
	//// 创建AgentSdk客户端，使用默认的Canister ID
	//client, err := AgentSdk.New(ag, "")
	//if err != nil {
	//	log.Fatalf("创建AgentSdk客户端失败: %v", err)
	//}
	//fmt.Println("创建AgentSdk客户端成功，使用身份:", principalID)
	//
	//// 保存密钥对到文件示例
	//fmt.Println("\n===== 示例5: 保存和加载密钥对到文件 =====")
	//keyFile := "example_keys.txt"
	//err = os.WriteFile(keyFile, []byte(pubKeyHex+"\n"+privKeyHex), 0600)
	//if err != nil {
	//	log.Printf("保存密钥到文件失败: %v", err)
	//} else {
	//	fmt.Printf("密钥对已保存到文件: %s\n", keyFile)
	//}
	//
	//// 从文件加载密钥对示例
	//fileContent, err := os.ReadFile(keyFile)
	//if err != nil {
	//	log.Printf("从文件加载密钥失败: %v", err)
	//} else {
	//	lines := strings.Split(string(fileContent), "\n")
	//	if len(lines) >= 2 {
	//		fmt.Printf("从文件加载的公钥: %s...\n", lines[0][:10])
	//		fmt.Printf("从文件加载的私钥: %s...\n", lines[1][:10])
	//	}
	//}
	//
	//// 删除测试文件
	//_ = os.Remove(keyFile)
	//
	//// 以下是原示例代码
	//fmt.Println("\n===== AgentSdk客户端功能演示 =====")
	//
	//// 查询代币余额示例
	//balance, err := client.GetBalance(principalID, "principal", "myTokenID")
	//if err != nil {
	//	log.Printf("查询余额失败: %v", err)
	//} else {
	//	fmt.Printf("余额: %s\n", (*balance).String())
	//}
	//
	//// 查询代币信息示例
	//tokenInfo, err := client.GetToken(principalID, "myTokenID")
	//if err != nil {
	//	log.Printf("查询代币信息失败: %v", err)
	//} else if tokenInfo.HasValue {
	//	token := tokenInfo.Value
	//	fmt.Printf("代币名称: %s, 符号: %s, 总供应量: %s\n",
	//		token.Name, token.Symbol, token.TotalSupply.String())
	//} else {
	//	fmt.Println("未找到代币信息")
	//}
	//
	//// 添加代币流动性示例
	//liquidityReq := AgentSdk.LiquidityRequest{
	//	TokenID:   "myTokenID",
	//	Amount:    big.NewInt(1000000),
	//	Operation: "add",
	//}
	//
	//liquidityResp, err := client.TokenLiquidity(liquidityReq)
	//if err != nil {
	//	log.Printf("添加流动性失败: %v", err)
	//} else if liquidityResp.OK != nil {
	//	fmt.Printf("添加流动性成功，操作ID: %s\n", (*liquidityResp.OK).String())
	//} else if liquidityResp.Err != nil {
	//	fmt.Printf("添加流动性失败，错误: %s\n", *liquidityResp.Err)
	//}
	//
	//// 交易代币示例
	//tradeReq := AgentSdk.TradeRequest{
	//	TokenID:   "myTokenID",
	//	Amount:    big.NewInt(500000),
	//	Operation: "buy",
	//}
	//
	//tradeResp, err := client.TokenTrade(tradeReq)
	//if err != nil {
	//	log.Printf("交易代币失败: %v", err)
	//} else if tradeResp.OK != nil {
	//	fmt.Printf("交易代币成功，操作ID: %s\n", (*tradeResp.OK).String())
	//} else if tradeResp.Err != nil {
	//	fmt.Printf("交易代币失败，错误: %s\n", *tradeResp.Err)
	//}
	//
	//// 计算百分比差异
	//diff := AgentSdk.CalculatePercentDifference(100.0, 90.0)
	//fmt.Printf("百分比差异: %.2f%%\n", diff)
	//
	//// BTC转换示例
	//btc := AgentSdk.ConvertToBTC(1000000)
	//fmt.Printf("BTC值: %.6f\n", btc)
}
