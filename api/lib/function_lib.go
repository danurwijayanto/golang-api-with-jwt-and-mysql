package lib

func ConvertStockToZero(stock uint64) uint64 {
	if stock < 0 {
		return 0
	}
	return stock
}
