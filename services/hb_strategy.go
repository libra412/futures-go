package services

import (
	"fmt"
	"github.com/libra412/futures-go/huobi"
	"github.com/libra412/futures-go/models"
)

//
func Huobi() error {
	fmt.Println("火币")
	res := huobi.SwapMarketHistoryKline("BTC-USD", "1min", 3)
	fmt.Println(res.Data)
	return nil
}

// 均线
func JunXian() error {

	return nil
}

// PIN策略
func tuDingStrategy(klines []models.Kline) {

}

// 孕线
func PregnantStrategy(klines []models.Kline) {

}
