package mock

import (
	"time"

	types "github.com/raanfefu/go-flow-cash/types"
)

func MockInputEvent() *types.Event {

	mov := make([]types.Movements, 3)
	mov[0] = types.Movements{
		Date:           time.Date(2022, 7, 1, 0, 0, 0, 0, &time.Location{}),
		Amount:         float64(-500000),
		IndexationRate: float64(0),
		PassMonth:      int32(0),
		MType:          "C",
	}
	mov[1] = types.Movements{
		Date:           time.Date(2022, 7, 2, 0, 0, 0, 0, &time.Location{}),
		Amount:         float64(-100000),
		IndexationRate: float64(0),
		PassMonth:      int32(0),
		MType:          "C",
	}
	mov[2] = types.Movements{
		Date:           time.Date(2022, 7, 3, 0, 0, 0, 0, &time.Location{}),
		Amount:         float64(-600000),
		IndexationRate: float64(0),
		PassMonth:      int32(0),
		MType:          "C",
	}

	event := &types.Event{
		LeaseAgreementType: "FF",
		DateOfPurchase:     time.Date(2022, 7, 1, 0, 0, 0, 0, &time.Location{}),
		Start:              time.Date(2022, 9, 1, 0, 0, 0, 0, &time.Location{}),
		End:                time.Date(2027, 10, 1, 0, 0, 0, 0, &time.Location{}),
		FistPayment:        time.Date(2022, 9, 1, 0, 0, 0, 0, &time.Location{}),
		PaymentPeriod:      1,
		PaymentAmount:      10000,
		IndexationRates: types.IndexationRates{
			IndexationRateValue: 4,
		},
		IndexationPeriod: 12,
		Capital:          mov,
	}

	return event
}

type AContent struct {
	Type string
	Id   string
}
