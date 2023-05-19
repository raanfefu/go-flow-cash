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
	tir := finances.XIRR(movements)
	finances.CalcCurrentValue(&result, event.DateOfPurchase, tir)

	for i := 0; i < len(movements); i++ {

		fmt.Printf("DT: %s | MV: %.2f | CF: %.2f | CV:%.2f\n", movements[i].Date, movements[i].Amount, movements[i].CashFlow, movements[i].CurrentValue)
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
	fmt.Print(r.TIR)

}
