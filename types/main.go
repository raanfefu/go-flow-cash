package structevent

import "time"

type Event struct {
	Name             string    `validate:"required"`
	IndexationRate   int64     `validate:"min=1"`
	IndexationPeriod int64     `validate:"min=1"`
	Start            time.Time `validate:"required"`
	End              time.Time `validate:"required"`
	FistPayment      time.Time `validate:"required"`
	PaymentPeriod    int64     `validate:"min=1"`
	Amount           int       `validate:"required,min=1"`
}

type FlowResult struct {
	Name string `json:@"name"`
}
