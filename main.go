package main

import (
	"fmt"

	mock "github.com/raanfefu/go-flow-cash/mock"
	types "github.com/raanfefu/go-flow-cash/types"
	validations "github.com/raanfefu/go-flow-cash/validations"

	"gopkg.in/go-playground/validator.v9"
)

func Handler(event *types.Event) (*types.FlowResult, error) {

	difference := event.End.Sub(event.Start)
	months := int32(difference.Hours()/24/30) / 1
	frecuency := months / event.PaymentPeriod

	nextDateIndexation := event.Start.AddDate(0, int(event.IndexationPeriod), 0)

	newDate := event.Start.AddDate(0, 0, 0)
	amount := event.Amount
	movements := make([]*types.Movements, frecuency)
	indexationFactor, _ := retriveIndexationRateValue(event.IndexationRates)

	for f := 0; f < int(frecuency); f++ {
		pastMonth := f * int(event.PaymentPeriod)

		if newDate.After(nextDateIndexation) {
			nextDateIndexation = nextDateIndexation.AddDate(0, int(event.IndexationPeriod), 0)
			amount = float32(amount) * indexationFactor
		}
		movements[f] = &types.Movements{
			Amount:         amount,
			Date:           newDate,
			IndexationRate: indexationFactor,
			PassMonth:      int32(pastMonth),
		}

		newDate = newDate.AddDate(0, int(event.PaymentPeriod), 0)
	}
	for i := 0; i < int(frecuency); i++ {
		fmt.Printf("%s | %f | %f\n", movements[i].Date, movements[i].Amount, movements[i].IndexationRate)
	}

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
