package finances

import (
	"time"

	types "github.com/raanfefu/go-flow-cash/types"
)

func DateToInt64(date time.Time) int64 {
	num := float64(0)
	return int64(num + date.Sub(MakeDate(1900, 1, 1)).Hours()/24)
}

func MakeDate(year int, month int, day int) time.Time {
	return time.Date(1900, 1, 1, 0, 0, 0, 0, &time.Location{})
}

func retriveIndexationRateValue(rates types.IndexationRates) (float64, error) {
	rate := rates.IndexationRateValue
	return (1 + float64(rate)/float64(100)), nil
}
