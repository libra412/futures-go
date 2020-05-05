package talib

// 均数
func SMA(prices []float64, period int) []float64 {
	count := len(prices)
	if count < period {
		return nil
	}
	avg := make([]float64, count-period+1)
	k := 0
	for i := 1; i < count; i++ {
		prices[i] += prices[i-1]
		if i >= (period - 1) {
			if i > (period - 1) {
				prices[i] -= prices[i-period]
			}
			avg[k] = prices[i] / float64(period)
			k++
		}
	}
	return avg
}
