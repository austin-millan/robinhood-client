package robinhood

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	model "gitlab.com/brokerage-api/robinhood-openapi/openapi"
)

// OptionsOrderOpts encapsulates common Options order choices
type OptionsOrderOpts struct {
	Quantity    float64
	Price       float64
	Direction   model.Direction
	TimeInForce model.TimeInForce
	Type        model.ExecutionType
	Side        model.Side
	Trigger     model.Trigger
}

// OrderOptions places a new order for options
func (c *Client) OrderOptions(q *model.OptionInstrument, o OptionsOrderOpts) (json.RawMessage, error) {
	b := model.OptionOrderInput{
		Account:     c.Account.Url,
		Direction:   o.Direction,
		TimeInForce: o.TimeInForce,
		Legs: []model.Leg{{
			Option:         q.Url,
			RatioQuantity:  "1",
			Side:           o.Side,
			PositionEffect: "open",
		}},
		Trigger:  o.Trigger,
		Type:     o.Type,
		Quantity: fmt.Sprintf("%f", o.Quantity),
		Price:    fmt.Sprintf("%f", o.Price),
		RefId:    uuid.New().String(),
	}

	if o.Side != model.BUY {
		b.Legs[0].PositionEffect = "close"
	}

	bs, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", EPOptions+"orders/", bytes.NewReader(bs))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	var out json.RawMessage
	err = c.DoAndDecode(req, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetOptionsOrders returns all outstanding options orders
func (c *Client) GetOptionsOrders(ctx context.Context) (*[]model.OptionOrder, error) {
	rs := make([]model.OptionOrder, 0)

	var results model.GetOptionOrdersResponse

	err := c.GetAndDecode(EPOptions+"orders/", &results)
	if err != nil {
		return nil, err
	}

	rs = append(rs, results.Results...)
	pager := Pager{Next: results.Next, Previous: results.Previous}
	for pager.HasMore() {
		err := pager.GetNext(c, &results)
		if err != nil {
			return &rs, err
		}
		rs = append(rs, results.Results...)
		select {
		case <-ctx.Done():
			return &rs, nil
		default:
		}
	}

	return &rs, nil
}
