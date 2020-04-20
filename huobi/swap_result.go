package huobi

//
type BaseHistoryKlineResult struct {
	Vol    int     `json:"vol"`
	Close  float64 `json:"close"`
	Count  int     `json:"count"`
	High   float64 `json:"high"`
	Id     int64   `json:"id"`
	Low    float64 `json:"low"`
	Open   float64 `json:"open"`
	Amount float64 `json:"amount"`
}

//
type SwapHistoryKlineResult struct {
	Ch     string                   `json:"ch"`
	Data   []BaseHistoryKlineResult `json:"data"`
	Status string                   `json:"status"`
	Ts     int64                    `json:"ts"`
}
