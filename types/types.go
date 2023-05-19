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
	DateOfPurchase     time.Time       `validate:"required"`
	PaymentPeriod      int32           `validate:"min=1"`
	PaymentAmount      float64         `validate:"required,min=1"`
	Capital            []Movements
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
	MovementsResult []Movements
	TIR             float64
}

type Movements struct {
	Date           time.Time
	Amount         float64
	CashFlow       float64
	IndexationRate float64
	PassMonth      int32
	MType          string
	CurrentValue   float64
}
