package finances

import (
	"math"
	"time"

	types "github.com/raanfefu/go-flow-cash/types"
)

const TYPE_INDEXATION_FIXED_FIXED = "FF"
const TYPE_INDEXATION_FIXED_VARIABLE = "FV"
const TYPE_INDEXATION_VARIABLE = "V"

func MakeCashFlow(event *types.Event) (types.FlowResult, error) {

	result := &types.FlowResult{}

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
	result.SigmaCapital = sumCapital

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
			indexationFactor, _ := retriveIndexationRateValue(event.IndexationRates)
			amount = amount * indexationFactor
		}

		movements = append(movements, types.Movements{
			Date:           newDate,
			Amount:         amount,
			Capital:        0,
			IndexationRate: indexationFactor,
			PassMonth:      int32(pastMonth),
			CashFlow:       amount,
		})

		newDate = newDate.AddDate(0, int(event.PaymentPeriod), 0)
	}

	movements[len(movements)-1].CashFlow = (-1 * sumCapital) + movements[len(movements)-1].Amount

	result.MovementsResult = movements

	return *result, nil
}

func CalcCurrentValue(result *types.FlowResult, dateOfPurchase time.Time) (*types.FlowResult, error) {

	movements := result.MovementsResult
	tir := result.TIR

	for i := 0; i < len(movements); i++ {

		base := (float64(1) + tir)
		dateOfPurchase := DateToInt64(dateOfPurchase)    //int64(event.DateOfPurchase.Sub(time.Date(1900, 1, 1, 0, 0, 0, 0, &time.Location{})).Hours() / 24)
		dateOfMovement := DateToInt64(movements[i].Date) //int64(movements[i].Date.Sub(time.Date(1900, 1, 1, 0, 0, 0, 0, &time.Location{})).Hours() / 24)

		pow := float64(dateOfPurchase-dateOfMovement) / float64(360) * -1
		movements[i].CurrentValue = movements[i].CashFlow / math.Pow(base, float64(pow))

		if movements[i].CurrentValue > 0 {
			result.SigmaCurrentValue = result.SigmaCurrentValue + movements[i].CurrentValue
		}

	}

	return result, nil
}

/*

	Formula:
	( [i].currentValue / SigmaCurrentVaue)*(([i].movDate  - startDate)/360)

*/

func Duration(result *types.FlowResult, startDate time.Time) (*types.FlowResult, error) {
	movements := result.MovementsResult

	for i := 0; i < len(movements)-1; i++ {
		if movements[i].CurrentValue > 0 {
			duration := (movements[i].CurrentValue / result.SigmaCurrentValue) * ((DateToFloat64(movements[i].Date) - DateToFloat64(startDate)) / float64(360))
			movements[i].Duration = duration
			result.SigmaDuration = result.SigmaDuration + duration
		}
	}
	last := len(movements) - 1
	duration := (movements[last].Amount / result.SigmaCurrentValue) * ((DateToFloat64(movements[last].Date) - DateToFloat64(startDate)) / float64(360))
	movements[last].Duration = duration
	result.SigmaDuration = result.SigmaDuration + duration

	result.Duration = result.SigmaDuration / (1 + result.TIR)

	return result, nil
}
