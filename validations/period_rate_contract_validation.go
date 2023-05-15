package validations

import (
	types "github.com/raanfefu/go-flow-cash/types"
	"gopkg.in/go-playground/validator.v9"
)

func PeriodAndRateContractValidation(sl validator.StructLevel) {
	event := sl.Current().Interface().(types.Event)
	if event.IndexationPeriod < event.PaymentPeriod {
		sl.ReportError(event.Start, "PaymentPeriod", "PaymentPeriod", "\"PaymentPeriod gt  IndexationPeriod\"", "None")
	}
}


//6f08ac6c-18a9-45f6-af7e-978527c9778d
//1553e0bc-e27e-4e16-be7a-a61a8ff8ea7c
