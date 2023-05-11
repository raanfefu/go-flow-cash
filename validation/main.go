package validations

import (
	"github.com/go-playground/validator"
	types "github.com/raanfefu/go-flow-cash/types"
)

func DateContractValidation(sl validator.StructLevel) {
	event := sl.Current().Interface().(types.Event)
	if event.Start.After(event.End) {
		sl.ReportError(event.Start, "Start", "Start", "\"Start before End\"", "None")
	}

	if event.Start.After(event.FistPayment) {
		sl.ReportError(event.Start, "Start", "Start", "\"Start after FistPayment\"", "None")
	}
}

func PeriodAndRateContractValidation(sl validator.StructLevel) {
	event := sl.Current().Interface().(types.Event)
	if event.IndexationRate > event.PaymentPeriod {
		sl.ReportError(event.Start, "PaymentPeriod", "PaymentPeriod", "\"PaymentPeriod gt  IndexationPeriod\"", "None")
	}
}
