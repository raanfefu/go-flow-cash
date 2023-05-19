package finances

import (
	"math"
	"time"

	types "github.com/raanfefu/go-flow-cash/types"
)

func MakeCashFlow(event *types.Event) (types.FlowResult, error) {

	indexationFactor := float64(0)
	cMovSize := len(event.Capital)
	difference := event.End.Sub(event.Start)
	months := int32(difference.Hours()/24/30) / 1
	frecuency := (months / event.PaymentPeriod) + int32(cMovSize)

	nextDateIndexation := event.Start.AddDate(0, int(event.IndexationPeriod-1), 0)

	newDate := event.Start.AddDate(0, 0, 0)
	amount := event.PaymentAmount
	movements := make([]types.Movements, cMovSize)

	sumCapital := float64(0)

	// Add Capital Movements
	for f := 0; f < int(cMovSize); f++ {
		movements[f] = event.Capital[f]
		movements[f].CashFlow = event.Capital[f].Amount
		sumCapital += event.Capital[f].Amount
	}

	// Calc Futures
	count := 0
	for f := cMovSize; f < int(frecuency); f++ {

		pastMonth := f * int(event.PaymentPeriod)
		if newDate.After(event.End) {
			break
		}
		count++
		if newDate.After(nextDateIndexation) {
			nextDateIndexation = nextDateIndexation.AddDate(0, int(event.IndexationPeriod), 0)
			factor, _ := retriveIndexationRateValue(event.IndexationRates)
			amount = amount * factor
			indexationFactor = factor
		}

		movements = append(movements, types.Movements{
			Amount:         amount,
			Date:           newDate,
			IndexationRate: indexationFactor,
			PassMonth:      int32(pastMonth),
			MType:          "R",
			CashFlow:       amount,
		})

		newDate = newDate.AddDate(0, int(event.PaymentPeriod), 0)
	}

	movements[len(movements)-1].CashFlow = (-1 * sumCapital) + movements[len(movements)-1].Amount
	movements[len(movements)-1].MType = "A"

	result := &types.FlowResult{
		MovementsResult: movements,
	}

	return *result, nil
}

func CalcCurrentValue(result *types.FlowResult, dateOfPurchase time.Time, tir float64) (*types.FlowResult, error) {

	movements := result.MovementsResult

	for i := 0; i < len(movements); i++ {

		base := (float64(1) + tir)
		dateOfPurchase := DateToInt64(dateOfPurchase)    //int64(event.DateOfPurchase.Sub(time.Date(1900, 1, 1, 0, 0, 0, 0, &time.Location{})).Hours() / 24)
		dateOfMovement := DateToInt64(movements[i].Date) //int64(movements[i].Date.Sub(time.Date(1900, 1, 1, 0, 0, 0, 0, &time.Location{})).Hours() / 24)

		pow := float64(dateOfPurchase-dateOfMovement) / float64(360) * -1
		movements[i].CurrentValue = movements[i].CashFlow / math.Pow(base, float64(pow))

	}

	return result, nil
}
