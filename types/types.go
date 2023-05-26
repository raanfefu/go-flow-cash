package structevent

import (
	"time"
)

type CustomDate time.Time

type Event struct {
	Capital            []Movements
	IndexationRates    IndexationRates `json:"indexationRate"`                           // Indexation Rate Parameters
	LeaseAgreementType string          `validate:"required" 		json:"leaseAgreementType"` // Lease Agreement Type ( FF / FV / )
	IndexationPeriod   int32           `validate:"min=1" 			json:"indexationPeriod"`
	Start              CustomDate      `validate:"required" 		json:"start"`
	End                CustomDate      `validate:"required" 		json:"end"`
	FistPayment        CustomDate      `validate:"required" 		json:"fistPayment"`
	DateOfPurchase     CustomDate      `validate:"required" 		json:"dateOfPurchase"`
	PaymentPeriod      int32           `validate:"min=1"    		json:"paymentPeriod"`
	PaymentAmount      float64         `validate:"required,min=1"   json:"paymentAmount"`
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
	MovementsResult   []Movements
	TIR               float64
	SigmaCapital      float64
	SigmaCurrentValue float64
	SigmaDuration     float64
	Duration          float64
}

type Movements struct {
	Date           time.Time
	Amount         float64
	Capital        float64
	CashFlow       float64
	IndexationRate float64
	CurrentValue   float64
	Duration       float64
	PassMonth      int32
}
