package validations

import (
	types "github.com/raanfefu/go-flow-cash/types"
	"gopkg.in/go-playground/validator.v9"
)

func PeriodAndRateContractValidation(sl validator.StructLevel) {
	event := sl.Current().Interface().(types.Event)
	if event.IndexationRate < event.PaymentPeriod {
		sl.ReportError(event.Start, "PaymentPeriod", "PaymentPeriod", "\"PaymentPeriod gt  IndexationPeriod\"", "None")
	}
}
