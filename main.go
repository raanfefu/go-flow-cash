package main

import (
	"fmt"

	goxirr "github.com/maksim77/goxirr"
	mock "github.com/raanfefu/go-flow-cash/mock"
	types "github.com/raanfefu/go-flow-cash/types"
	validations "github.com/raanfefu/go-flow-cash/validations"
	"gopkg.in/go-playground/validator.v9"
)

func Handler(event *types.Event) (*types.FlowResult, error) {

	difference := event.End.Sub(event.Start)
	months := int32(difference.Hours()/24/30) / 1
	frecuency := months / event.PaymentPeriod

	nextDateIndexation := event.Start.AddDate(0, int(event.IndexationPeriod-1), 0)

	newDate := event.Start.AddDate(0, 0, 0)
	amount := event.PaymentAmount
	movements := make([]types.Movements, frecuency)
	tx := make(goxirr.Transactions, frecuency)

	indexationFactor := float32(0)

	for f := 0; f < int(frecuency); f++ {
		pastMonth := f * int(event.PaymentPeriod)

		if newDate.After(nextDateIndexation) {
			nextDateIndexation = nextDateIndexation.AddDate(0, int(event.IndexationPeriod), 0)
			factor, _ := retriveIndexationRateValue(event.IndexationRates)
			amount = float32(amount) * float32(factor)
			indexationFactor = factor
		}
		tx[f+3] = goxirr.Transaction{
			Date: newDate,
			Cash: float64(amount),
		}
		movements[f] = types.Movements{
			Amount:         amount,
			Date:           newDate,
			IndexationRate: indexationFactor,
			PassMonth:      int32(pastMonth),
		}

		newDate = newDate.AddDate(0, int(event.PaymentPeriod), 0)
	}

	fmt.Println(goxirr.Xirr(tx))
	/*for i := 0; i < int(frecuency); i++ {
		fmt.Printf("%s | %f | %f\n",
			movements[i].Date.Format("2006-01-02"),
			movements[i].Amount,
			movements[i].IndexationRate,
		)
	}*/

	result := &types.FlowResult{}
	return result, nil
}

func retriveIndexationRateValue(rates types.IndexationRates) (float32, error) {
	rate := rates.IndexationRateValue
	return (1 + float32(rate)/float32(100)), nil
}

func main() {

	validate := validator.New()
	validate.RegisterStructValidation(validations.DateContractValidation, &types.Event{})
	validate.RegisterStructValidation(validations.PeriodAndRateContractValidation, &types.Event{})

	event := mock.MockInputEvent()

	err := validate.Struct(event)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field(), err.Tag())
		}
		return
	}
	Handler(event)
}
