package finances

import (
	"math"

	types "github.com/raanfefu/go-flow-cash/types"
)

const MAX_LEVEL = 1
const INIT_LEVEL = 0
const INIT_GUESS = 0

func XIRR(transactions *[]types.Movements) float64 {
	tir := iteration(transactions, INIT_GUESS, INIT_LEVEL, MAX_LEVEL)
	return tir
}

func iteration(transactions *[]types.Movements, ini float64, level int, maxLevel int) float64 {
	level += 1
	tir := float64(0)
	for iw := 0; iw < 99; iw++ {
		guess := (ini + (float64(iw))/math.Pow(float64(100), float64(level)))
		sigma := sigmaResolve(*transactions, guess)
		if sigma < 0 {
			//fmt.Printf("level: %d, it: %d\n", level, iw-1)
			if level < 100 {
				tir = iteration(transactions, tir, level, maxLevel)
			}

			break
		}
		tir = guess
	}
	return tir
}

func sigmaResolve(transactions []types.Movements, guess float64) float64 {
	initialDate := DateToInt64(transactions[0].Date)
	sum := float64(0)

	for i := 0; i < len(transactions); i++ {

		numDate := DateToInt64(transactions[i].Date) //int64(num + transactions[i].Date.Sub(time.Date(1900, 1, 1, 0, 0, 0, 0, &time.Location{})).Hours()/24)
		q := float64(numDate - initialDate)
		k := math.Round((float64(q)/float64(365.00000))*100000) / 100000
		w := math.Round((math.Pow(float64(1+guess), float64(k)))*100000000) / float64(100000000)
		terms := math.Round(float64(transactions[i].CashFlow/w)*10) / float64(10)
		sum += terms

	}
	return sum
}
