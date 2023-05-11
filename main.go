package main

import (
	"fmt"
	"time"

	types "github.com/raanfefu/go-flow-cash/types"
	validations "github.com/raanfefu/go-flow-cash/validations"

	strftime "github.com/itchyny/timefmt-go"
	"gopkg.in/go-playground/validator.v9"
)

func Handler(event *types.Event) (*types.FlowResult, error) {

	difference := event.End.Sub(event.Start)
	months := int64(difference.Hours()/24/30) / 1
	frecuency := int64(months / event.PaymentPeriod)
	numPeriods := int64(months / event.IndexationPeriod)

	fmt.Printf("num months %d\n", months)
	fmt.Printf("num payments %d\n", frecuency)
	fmt.Printf("num  indexation period %d\n", numPeriods)

	newDate := event.Start.AddDate(0, 0, 0)

	for f := 1; f <= int(frecuency); f++ {
		fmt.Printf("payment #%d - %s \n", f, strftime.Format(newDate, "%Y-%m-%d"))
		newDate = newDate.AddDate(0, int(event.PaymentPeriod), 0)
	}

	result := &types.FlowResult{Name: event.Name}
	return result, nil
}

func main() {
	validate := validator.New()
	validate.RegisterStructValidation(validations.DateContractValidation, &types.Event{})
	validate.RegisterStructValidation(validations.PeriodAndRateContractValidation, &types.Event{})
	event := mock()
	err := validate.Struct(event)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field(), err.Tag())
		}
		return
	}
	Handler(event)
}

func mock() *types.Event {
	event := &types.Event{
		Name:             "Data",
		Start:            time.Date(2022, 9, 1, 0, 0, 0, 0, &time.Location{}),
		End:              time.Date(2023, 9, 1, 0, 0, 0, 0, &time.Location{}),
		FistPayment:      time.Date(2022, 9, 1, 0, 0, 0, 0, &time.Location{}),
		PaymentPeriod:    1,
		Amount:           100000,
		IndexationRate:   1,
		IndexationPeriod: 1,
	}
	return event
}
