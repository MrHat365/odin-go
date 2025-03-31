package services

import (
	"encoding/json"
	"fmt"

	"github.com/MrHat365/odin-go/odin_api/client"
	"github.com/MrHat365/odin-go/odin_api/models"
)

// GetBTCPrice 获取比特币当前价格信息
// 返回比特币价格和相关信息
func GetBTCPrice() (*models.BTCInfo, error) {
	c := client.NewClient()

	// 发送请求
	endpoint := "/currency/btc"
	resp, err := c.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("获取比特币价格失败: %w", err)
	}

	// 解析响应
	var btcInfo models.BTCInfo
	if err := json.Unmarshal(resp, &btcInfo); err != nil {
		return nil, fmt.Errorf("解析比特币价格信息失败: %w", err)
	}

	return &btcInfo, nil
}
