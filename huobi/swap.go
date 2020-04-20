package huobi

import (
	"encoding/json"
	"strconv"
)

/*
获取K线
GET /swap-ex/market/history/kline
*/
func SwapMarketHistoryKline(strSymbol, strPeriod string, nSize int) SwapHistoryKlineResult {

	mapParams := make(map[string]string)
	mapParams["contract_code"] = strSymbol
	mapParams["period"] = strPeriod
	mapParams["size"] = strconv.Itoa(nSize)

	strRequestUrl := "/swap-ex/market/history/kline"
	strUrl := SWAP_URL + strRequestUrl

	jsonKLineReturn := HttpGetRequest(strUrl, mapParams)
	var v SwapHistoryKlineResult
	json.Unmarshal([]byte(jsonKLineReturn), &v)
	return v
}
