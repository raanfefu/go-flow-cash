package mock

import (
	"time"

	types "github.com/raanfefu/go-flow-cash/types"
)

func MockInputEvent() *types.Event {
	a := []types.Movements{{Type: "Hello", Id: "asdf"}, {Type: "World", Id: "ghij"}}

	event := &types.Event{
		LeaseAgreementType: "FF",
		Start:              time.Date(2022, 9, 1, 0, 0, 0, 0, &time.Location{}),
		End:                time.Date(2025, 9, 1, 0, 0, 0, 0, &time.Location{}),
		FistPayment:        time.Date(2022, 9, 1, 0, 0, 0, 0, &time.Location{}),
		PaymentPeriod:      1,
		Amount:             10000,
		IndexationRates: types.IndexationRates{
			IndexationRateValue: 4,
		},
		IndexationPeriod: 12,
		capital:          "a",
	}

	return event
}

type AContent struct {
	Type string
	Id   string
}
