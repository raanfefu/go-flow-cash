package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-playground/validator"
	"github.com/raanfefu/go-flow-cash/finances"
	types "github.com/raanfefu/go-flow-cash/types"
)

func (c *types.CustomDate) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse("2006-01-02", value) //parse time
	if err != nil {
		return err
	}
	*c = types.CustomDate(t) //set result using the pointer
	return nil
}

func (c types.CustomDate) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(c).Format("2006-01-02") + `"`), nil
}

func HandleRequest(ctx context.Context, event *types.Event) (*types.FlowResult, error) {

	result, _ := finances.MakeCashFlow(event)
	movements := result.MovementsResult
	result.TIR = finances.XIRR(&movements)
	finances.CalcCurrentValue(&result, event.DateOfPurchase)
	finances.Duration(&result, event.Start)
	return &result, nil
}

func ErrorsLambda(err error) error {
	for _, err := range err.(validator.ValidationErrors) {
		fmt.Println(err.Field(), err.Tag())
	}
	return err
}

func HandleRequestTest(ctx context.Context, event *types.Event) (*types.Event, error) {
	validate := validator.New()

	err := validate.Struct(event)
	if err != nil {
		return nil, ErrorsLambda(err)
	}

	return event, nil
}
func main() {

	lambda.Start(HandleRequestTest)

}
