package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/libra412/futures-go/models"
	"github.com/libra412/futures-go/okex"
)

// 创建okex平台请求器
func NewOKExClient() *okex.Client {
	var config okex.Config
	config.Endpoint = "https://www.okex.me/"
	config.ApiKey = ""
	config.SecretKey = ""
	config.Passphrase = ""
	config.TimeoutSecond = 45
	config.IsPrint = false
	config.I18n = okex.ENGLISH

	client := okex.NewClient(config)
	return client
}

var opName = map[string]string{"md": "买多", "mk": "买空"}

// 4小时策略
func F4HourUsdt() error {
	// 获取K线数据
	c := NewOKExClient()
	m := map[string]string{}
	m["start"] = time.Now().Add(-time.Hour*2*4).UTC().Format("2006-01-02T15:04:05") + ".000Z"
	m["granularity"] = "14400" // 14400 = 4h
	res, err := c.GetSwapCandlesByInstrument("BTC-USDT-SWAP", m)
	if err == nil { // 获取K线数据， 每次获取 前两条k线
		array := *res
		count := len(array)
		klines := make([]models.Kline, count)
		for i := 0; i < count; i++ {
			kline := models.Kline{}
			kline.Timestamp = array[i][0].(string)
			kline.Open = array[i][1].(string)
			kline.High = array[i][2].(string)
			kline.Low = array[i][3].(string)
			kline.Close = array[i][4].(string)
			kline.Volume = array[i][5].(string)
			kline.CurrencyVolume = array[i][6].(string)
			kline.Type = 14400
			kline.CreateTime = time.Now()
			klines[i] = kline
		}
		op, stop := strategy(klines)
		fmt.Println(time.Now().Local().Format("2006-01-02 15:04:05"), op, stop, "BTC-USDT-SWAP")
		if stop > 0 {
			// go sendStrategy("4小时USDT策略", op, strconv.FormatFloat(stop, 'f', 4, 64), "BTC-USDT-SWAP")
			dingdingRobot("4小时USDT策略:" + time.Now().Local().Format("2006-01-02 15:04:05") + ",策略:" + opName[op] + ",止损位:" + strconv.FormatFloat(stop, 'f', 4, 64))
		}
	}
	return nil
}

// 1小时策略
func F1HourUsdt() error {
	// 获取K线数据
	c := NewOKExClient()
	m := map[string]string{}
	m["start"] = time.Now().Add(-time.Hour*2).UTC().Format("2006-01-02T15:04:05") + ".000Z"
	m["granularity"] = "3600" // 3600s = 1h
	res, err := c.GetSwapCandlesByInstrument("BTC-USDT-SWAP", m)
	if err == nil { // 获取K线数据， 每次获取 前两条k线
		array := *res
		count := len(array)
		klines := make([]models.Kline, count)
		for i := 0; i < count; i++ {
			kline := models.Kline{}
			kline.Timestamp = array[i][0].(string)
			kline.Open = array[i][1].(string)
			kline.High = array[i][2].(string)
			kline.Low = array[i][3].(string)
			kline.Close = array[i][4].(string)
			kline.Volume = array[i][5].(string)
			kline.CurrencyVolume = array[i][6].(string)
			kline.Type = 3600
			kline.CreateTime = time.Now()
			klines[i] = kline
		}
		op, stop := strategy(klines)
		fmt.Println(time.Now().Local().Format("2006-01-02 15:04:05"), op, stop, "BTC-USDT-SWAP")
		if stop > 0 {
			go sendStrategy("1小时USDT策略", op, strconv.FormatFloat(stop, 'f', 4, 64), "BTC-USDT-SWAP")
			dingdingRobot("1小时USDT策略:" + time.Now().Local().Format("2006-01-02 15:04:05") + ",策略:" + opName[op] + ",止损位:" + strconv.FormatFloat(stop, 'f', 4, 64))
		}
	}
	return nil
}

// 4小时策略
func F4HourUsd() error {
	// 获取K线数据
	c := NewOKExClient()
	m := map[string]string{}
	m["start"] = time.Now().Add(-time.Hour*2*4).UTC().Format("2006-01-02T15:04:05") + ".000Z"
	m["granularity"] = "14400" // 14400 = 4h
	res, err := c.GetSwapCandlesByInstrument("BTC-USD-SWAP", m)
	if err == nil { // 获取K线数据， 每次获取 前两条k线
		array := *res
		count := len(array)
		klines := make([]models.Kline, count)
		for i := 0; i < count; i++ {
			kline := models.Kline{}
			kline.Timestamp = array[i][0].(string)
			kline.Open = array[i][1].(string)
			kline.High = array[i][2].(string)
			kline.Low = array[i][3].(string)
			kline.Close = array[i][4].(string)
			kline.Volume = array[i][5].(string)
			kline.CurrencyVolume = array[i][6].(string)
			kline.Type = 14400
			kline.CreateTime = time.Now()
			klines[i] = kline
		}
		op, stop := strategy(klines)
		fmt.Println(time.Now().Local().Format("2006-01-02 15:04:05"), op, stop, "BTC-USD-SWAP")
		if stop > 0 {
			// go sendStrategy("4小时USD策略", op, strconv.FormatFloat(stop, 'f', 4, 64), "BTC-USD-SWAP")
			dingdingRobot("4小时USD策略:" + time.Now().Local().Format("2006-01-02 15:04:05") + ",策略:" + opName[op] + ",止损位:" + strconv.FormatFloat(stop, 'f', 4, 64))
		}
	}
	return nil
}

// 1小时策略
func F1HourUsd() error {
	// 获取K线数据
	c := NewOKExClient()
	m := map[string]string{}
	m["start"] = time.Now().Add(-time.Hour*2).UTC().Format("2006-01-02T15:04:05") + ".000Z"
	m["granularity"] = "3600" // 3600s = 1h
	res, err := c.GetSwapCandlesByInstrument("BTC-USD-SWAP", m)
	if err == nil { // 获取K线数据， 每次获取 前两条k线
		array := *res
		count := len(array)
		klines := make([]models.Kline, count)
		for i := 0; i < count; i++ {
			kline := models.Kline{}
			kline.Timestamp = array[i][0].(string)
			kline.Open = array[i][1].(string)
			kline.High = array[i][2].(string)
			kline.Low = array[i][3].(string)
			kline.Close = array[i][4].(string)
			kline.Volume = array[i][5].(string)
			kline.CurrencyVolume = array[i][6].(string)
			kline.Type = 3600
			kline.CreateTime = time.Now()
			klines[i] = kline
		}
		op, stop := strategy(klines)
		fmt.Println(time.Now().Local().Format("2006-01-02 15:04:05"), op, stop, "BTC-USD-SWAP")
		if stop > 0 {
			go sendStrategy("1小时USD策略", op, strconv.FormatFloat(stop, 'f', 4, 64), "BTC-USD-SWAP")
			dingdingRobot("1小时USD策略:" + time.Now().Local().Format("2006-01-02 15:04:05") + ",策略:" + opName[op] + ",止损位:" + strconv.FormatFloat(stop, 'f', 4, 64))
		}
	}
	return nil
}

// 过滤器 被吞没的线不大于50 影线也要包
func filter1(openBarCurrent, closeBarCurrent, openBarPrevious, closeBarPrevious, highBarCurrent, lowBarCurrent, highBarPrevious, lowBarPrevious float64) bool {
	if math.Abs(openBarPrevious-closeBarPrevious) < 15 {
		if (highBarCurrent > highBarPrevious) && (lowBarCurrent < lowBarPrevious) {
			return true
		}
		return false
	}
	return true
}

// 过滤器 完全相等的实体  做空上影线要包 做多下影线要包
func filter2(openBarCurrent, closeBarCurrent, openBarPrevious, closeBarPrevious, highBarCurrent, lowBarCurrent, highBarPrevious, lowBarPrevious float64) bool {
	if closeBarCurrent == openBarPrevious { // 当前的收盘价等于之前的开盘价 == 等实体
		if openBarCurrent > closeBarCurrent { // 跌，做空的情况
			return highBarCurrent > highBarPrevious
		} else { // 涨，做多的情况
			return lowBarCurrent < lowBarPrevious
		}
	}
	return true
}

// 过滤器 被吞没的线不大于50 做空上影线要包 做多下影线要包
func filter3(openBarCurrent, closeBarCurrent, openBarPrevious, closeBarPrevious, highBarCurrent, lowBarCurrent, highBarPrevious, lowBarPrevious float64) bool {
	if math.Abs(openBarPrevious-closeBarPrevious) < 15 {
		if openBarCurrent > closeBarCurrent { // 跌，做空的情况
			return highBarCurrent >= highBarPrevious
		} else { // 涨，做多的情况
			return lowBarCurrent <= lowBarPrevious
		}
	}
	return true
}

// 吞没形态策略
func strategy(klines []models.Kline) (string, float64) {
	if len(klines) > 1 {
		var openBarCurrent, closeBarCurrent, openBarPrevious, closeBarPrevious, highBarCurrent, lowBarCurrent, highBarPrevious, lowBarPrevious float64
		if len(klines) == 2 {
			openBarCurrent, _ = strconv.ParseFloat(klines[0].Open, 64)
			closeBarCurrent, _ = strconv.ParseFloat(klines[0].Close, 64)
			openBarPrevious, _ = strconv.ParseFloat(klines[1].Open, 64)
			closeBarPrevious, _ = strconv.ParseFloat(klines[1].Close, 64)
			highBarCurrent, _ = strconv.ParseFloat(klines[0].High, 64)
			lowBarCurrent, _ = strconv.ParseFloat(klines[0].Low, 64)
			highBarPrevious, _ = strconv.ParseFloat(klines[1].High, 64)
			lowBarPrevious, _ = strconv.ParseFloat(klines[1].Low, 64)
		} else if len(klines) == 3 {
			openBarCurrent, _ = strconv.ParseFloat(klines[1].Open, 64)
			closeBarCurrent, _ = strconv.ParseFloat(klines[1].Close, 64)
			openBarPrevious, _ = strconv.ParseFloat(klines[2].Open, 64)
			closeBarPrevious, _ = strconv.ParseFloat(klines[2].Close, 64)
			highBarCurrent, _ = strconv.ParseFloat(klines[1].High, 64)
			lowBarCurrent, _ = strconv.ParseFloat(klines[1].Low, 64)
			highBarPrevious, _ = strconv.ParseFloat(klines[2].High, 64)
			lowBarPrevious, _ = strconv.ParseFloat(klines[2].Low, 64)
		}
		// 做多 默认openBarCurrent <= closeBarPrevious成立
		if (openBarCurrent < openBarPrevious) && (closeBarCurrent > openBarPrevious) {
			stopDuo := openBarCurrent
			if lowBarCurrent < lowBarPrevious {
				stopDuo = lowBarCurrent - 10
			} else {
				stopDuo = lowBarPrevious - 10
			}
			// 过滤器1
			res1 := filter3(openBarCurrent, closeBarCurrent, openBarPrevious, closeBarPrevious, highBarCurrent, lowBarCurrent, highBarPrevious, lowBarPrevious)
			if !res1 {
				return "过滤器1不做多单", 0
			}
			// 过滤器2
			res2 := filter2(openBarCurrent, closeBarCurrent, openBarPrevious, closeBarPrevious, highBarCurrent, lowBarCurrent, highBarPrevious, lowBarPrevious)
			if !res2 {
				return "过滤器2不做多单", 0
			}
			return "md", stopDuo
		}
		// 做空 默认openBarCurrent >= closeBarPrevious
		if (openBarCurrent > openBarPrevious) && (closeBarCurrent < openBarPrevious) {
			stopKong := openBarCurrent
			if highBarCurrent > highBarPrevious {
				stopKong = highBarCurrent + 10
			} else {
				stopKong = highBarPrevious + 10
			}
			// 过滤器1
			res1 := filter3(openBarCurrent, closeBarCurrent, openBarPrevious, closeBarPrevious, highBarCurrent, lowBarCurrent, highBarPrevious, lowBarPrevious)
			if !res1 {
				return "过滤器1不做空单", 0
			}
			// 过滤器2
			res2 := filter2(openBarCurrent, closeBarCurrent, openBarPrevious, closeBarPrevious, highBarCurrent, lowBarCurrent, highBarPrevious, lowBarPrevious)
			if !res2 {
				return "过滤器2不做空单", 0
			}
			return "mk", stopKong
		}
	}
	return "不做单", 0
}

// 钉钉机器人推送
func dingdingRobot(content string) {
	requestUrl := `https://oapi.dingtalk.com/robot/send?access_token=8beb7ff4e5231f159e98a1f5b6e331ad6a79099b233d27e2c702dd84bb4ff8a2`
	jsonStr := `{"msgtype": "text", 
        "text": {
             "content": "` + content + `"
        }
      }`
	req, _ := http.NewRequest("POST", requestUrl, strings.NewReader(jsonStr))
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("推送出错", err)
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("推送结果", string(body))
}

// 发送指令
func sendStrategy(name, direction, stopPoint, instrument string) {
	reqUrl := "http://127.0.0.1:4000/btc/acceptStrategyReal"
	params := map[string]string{"Name": name, "Direction": direction,
		"StopPoint": stopPoint, "Instrument": instrument}
	reqbody, _ := json.Marshal(params)

	req, _ := http.NewRequest("POST", reqUrl, strings.NewReader(string(reqbody)))
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("发送出错", err)
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("发送结果", string(body))
}
