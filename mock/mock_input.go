package mock

import (
	"time"

	types "github.com/raanfefu/go-flow-cash/types"
)

func MockInputEvent() *types.Event {
	event := &types.Event{
		Name:             "Data",
		Start:            time.Date(2022, 9, 1, 0, 0, 0, 0, &time.Location{}),
		End:              time.Date(2025, 9, 1, 0, 0, 0, 0, &time.Location{}),
		FistPayment:      time.Date(2022, 9, 1, 0, 0, 0, 0, &time.Location{}),
		PaymentPeriod:    1,
		Amount:           10000,
		IndexationRate:   4,
		IndexationPeriod: 12,
	}
	return event
}
