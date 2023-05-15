package mock

import (
	"time"

	types "github.com/raanfefu/go-flow-cash/types"
)

func MockInputEvent() *types.Event {

	mov := make([]types.Movements, 4)
	mov[0] = types.Movements{
		Date:           time.Date(2022, 7, 1, 0, 0, 0, 0, &time.Location{}),
		Amount:         float32(-500),
		IndexationRate: float32(0),
		PassMonth:      int32(0),
		mType:          "C",
	}

	event := &types.Event{
		LeaseAgreementType: "FF",
		Start:              time.Date(2022, 9, 1, 0, 0, 0, 0, &time.Location{}),
		End:                time.Date(2025, 9, 1, 0, 0, 0, 0, &time.Location{}),
		FistPayment:        time.Date(2022, 9, 1, 0, 0, 0, 0, &time.Location{}),
		PaymentPeriod:      1,
		PaymentAmount:      10000,
		IndexationRates: types.IndexationRates{
			IndexationRateValue: 4,
		},
		IndexationPeriod: 3,
		Capital:          mov,
	}

	return event
}

type AContent struct {
	Type string
	Id   string
}
