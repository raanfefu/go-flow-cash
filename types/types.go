package structevent

import (
	"time"
)

type Event struct {
	IndexationRates    IndexationRates `validate:"required"`
	LeaseAgreementType string          `validate:"required"`
	IndexationPeriod   int32           `validate:"min=1"`
	Start              time.Time       `validate:"required"`
	End                time.Time       `validate:"required"`
	FistPayment        time.Time       `validate:"required"`
	PaymentPeriod      int32           `validate:"min=1"`
	PaymentAmount      float32         `validate:"required,min=1"`
	//ContratAmount      float32         `validate:"required,min=1"`
	Capital []Movements
}

type IndexationRates struct {
	Rates               []IndexationRate
	IndexationRateValue int32
}

type IndexationRate struct {
	IndexationRateValue int32
	Date                time.Time
}

type FlowResult struct {
	Name string `json:@"name"`
}

type Movements struct {
	Date           time.Time
	Amount         float32
	IndexationRate float32
	PassMonth      int32
	mType          string
}
