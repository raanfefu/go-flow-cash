package mock

import (
	"time"

	types "github.com/raanfefu/go-flow-cash/types"
)

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
