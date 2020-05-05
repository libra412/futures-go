package services

import (
	"github.com/libra412/futures-go/models"
	"github.com/libra412/futures-go/talib"
)

// 双均线策略  fast短周期，slow长周期
func BisexualStrategy(klines []models.Kline, fast, slow int) string {
	prices := make([]float64, len(klines))
	for i := 0; i < len(klines); i++ {
		prices[i] = klines[i].Close.(float64)
	}
	fastAvg := talib.SMA(prices, fast)
	slowAvg := talib.SMA(prices, slow)

	// slow的上一个点 小于 fast的上一个点 并且 slow的当前这个点 大于等于 fast的当前这个点 为 均线下穿，做空
	if slowAvg[len(slowAvg)-2] < fastAvg[len(fastAvg)-2] && slowAvg[len(slowAvg)-1] >= fastAvg[len(fastAvg)-1] {
		// 平多仓
		// 开空仓
		return "mk"
	}
	// fast的上一个点 小于 slow的上一个点 并且 fast的当前这个点 大于等于 slow的当前这个点 为 均线上穿，做多
	if slowAvg[len(slowAvg)-2] > fastAvg[len(fastAvg)-2] && slowAvg[len(slowAvg)-1] <= fastAvg[len(fastAvg)-1] {
		// 平空仓
		// 开多仓
		return "md"
	}
	return "none"
}

// 网格策略
func name() {

}
