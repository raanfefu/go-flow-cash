package validations

import (
	types "github.com/raanfefu/go-flow-cash/types"
	"gopkg.in/go-playground/validator.v9"
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
