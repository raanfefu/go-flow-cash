package main

import (
	"fmt"

	finances "github.com/raanfefu/go-flow-cash/finances"
	mock "github.com/raanfefu/go-flow-cash/mock"
	types "github.com/raanfefu/go-flow-cash/types"
	validations "github.com/raanfefu/go-flow-cash/validations"
	"gopkg.in/go-playground/validator.v9"
)

func Handler(event *types.Event) (*types.FlowResult, error) {

	result, _ := finances.MakeCashFlow(event)
	movements := result.MovementsResult
	result.TIR = finances.XIRR(&movements)
	finances.CalcCurrentValue(&result, event.DateOfPurchase)
	finances.Duration(&result, event.Start)
	for i := 0; i < len(movements); i++ {

		fmt.Printf("DT: %s | MV: %.2f | CF: %.2f | CV:%.2f | DU: %.6f \n", movements[i].Date, movements[i].Amount, movements[i].CashFlow, movements[i].CurrentValue, movements[i].Duration)
	}
	return &result, nil
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
	r, _ := Handler(event)
	fmt.Printf("TIR: %f.2\n", r.TIR)
	fmt.Printf("Sigma Capital: %.2f\n", r.SigmaCapital)
	fmt.Printf("Sigma CurrentValue: %.2f\n", r.SigmaCurrentValue)
	fmt.Printf("Sigma Duration: %.2f\n", r.SigmaDuration)
	fmt.Printf("Duration: %.2f\n ", r.Duration)

}
