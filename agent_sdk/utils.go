package agent_sdk

import (
	"math"
	"math/big"
)

// ConvertToBTC 将satoshis转换为BTC
// 与C#版本对应，将satoshis值除以1000，并保留6位小数
func ConvertToBTC(satoshis int64) float64 {
	// 转换为浮点数后除以1000
	value := float64(satoshis) / 1000.0
	// 四舍五入到6位小数
	return math.Round(value*1000000) / 1000000
}

// ConvertToTokenAmount 将satoshis转换为代币数量
// 将satoshis值除以100000000000并四舍五入
func ConvertToTokenAmount(satoshis int64) int64 {
	// 转换为浮点数后除以100000000000
	value := float64(satoshis) / 100000000000.0
	// 四舍五入到整数
	return int64(math.Round(value))
}

// CalculatePercentDifference 计算两个数值之间的百分比差异
// 计算公式: |A - B| / ((A + B) / 2) * 100
func CalculatePercentDifference(a, b float64) float64 {
	difference := math.Abs(a - b)
	average := (a + b) / 2
	return (difference / average) * 100
}

// NewTokenAmount 从int64创建TokenAmount类型
func NewTokenAmount(value int64) *big.Int {
	return big.NewInt(value)
}

// NewTokenAmountFromString 从字符串创建TokenAmount类型
func NewTokenAmountFromString(value string) (*big.Int, bool) {
	n := new(big.Int)
	success := true
	n, success = n.SetString(value, 10)
	if !success {
		return big.NewInt(0), false
	}
	return n, true
}

// TokenAmountToFloat64 将TokenAmount转换为float64类型，用于显示
func TokenAmountToFloat64(amount *big.Int) float64 {
	if amount == nil {
		return 0
	}

	// 如果数值太大无法表示为float64，则取一个近似值
	// 这可能导致精度损失，但对于显示目的通常足够
	f := new(big.Float).SetInt(amount)
	result, _ := f.Float64()
	return result
}
