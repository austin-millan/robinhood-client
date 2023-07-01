package robinhood

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	model "gitlab.com/brokerage-api/robinhood-client/models"
)

// OrderSide is which side of the trade an order is on
type OrderSide int

// OrderOpts encapsulates differences between order types
type OrderOpts struct {
	Side          model.Side
	Type          model.ExecutionType
	Quantity      uint64
	Price         float64
	TimeInForce   model.TimeInForce
	ExtendedHours bool
	Stop, Force   bool
}

// Order places an order for a given instrument
func (c *Client) Order(i *model.InstrumentData, o OrderOpts) (*model.Order, error) {
	a := model.Order{
		Account:       c.Account.Url,
		Instrument:    i.Url,
		Symbol:        i.Symbol,
		Type:          o.Type,
		TimeInForce:   o.TimeInForce,
		Quantity:      string(o.Quantity),
		Side:          o.Side,
		ExtendedHours: o.ExtendedHours,
		Price:         fmt.Sprintf("%f", o.Price),
		Trigger:       "immediate",
	}

	if o.Stop {
		a.StopPrice = fmt.Sprintf("%f", o.Price)
		a.Trigger = "stop"
	}

	bs, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}

	post, err := http.NewRequest("POST", EPOrders, bytes.NewReader(bs))
	post.Header.Add("Content-Type", "application/json")

	var out model.Order
	err = c.DoAndDecode(post, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

// UpdateOrder returns any errors and updates the item with any recent changes.
func (c *Client) UpdateOrder(o model.Order) error {
	return c.GetAndDecode(o.Url, o)
}

// CancelOrder attempts to cancel an odrer
func (c *Client) CancelOrder(o *model.Order) error {
	post, err := http.NewRequest("POST", o.CancelUrl, nil)
	if err != nil {
		return err
	}

	var o2 model.Order
	err = c.DoAndDecode(post, &o2)
	if err != nil {
		return errors.Wrap(err, "could not decode response")
	}

	if o2.RejectReason != "" {
		return errors.New(o2.RejectReason)
	}
	return nil
}

// GetStockOrders returns orders made by this client.
func (c *Client) GetStockOrders() ([]model.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	rs := make([]model.Order, 0)
	defer cancel()

	var results model.PaginatedOrder
	// err := c.GetAndDecode(EPBase+"/orders?page_size=200", &results)
	err := c.GetAndDecode(EPOrders, &results)
	if err != nil {
		return nil, err
	}

	rs = append(rs, results.Results...)
	pager := Pager{Next: results.Next, Previous: results.Previous}
	for pager.HasMore() {
		err := pager.GetNext(c, &results)
		if err != nil {
			return rs, err
		}
		rs = append(rs, results.Results...)

		select {
		case <-ctx.Done():
			return rs, nil
		default:
		}
	}

	for _, order := range rs {  // TODO: optimize
		instrumentData, err := c.GetInstrument(order.Instrument)
		if err != nil {
			break
		}
		order.Symbol = instrumentData.Symbol
	}

	return rs, nil
}

// RecentOrders returns any recent orders made by this client.
func (c *Client) RecentOrders() ([]model.Order, error) {
	var o struct {
		Results []model.Order
	}
	err := c.GetAndDecode(EPOrders, &o)
	if err != nil {
		return nil, err
	}
	for _, order := range o.Results {
		instrumentData, err := c.GetInstrument(order.Instrument)
		if err != nil {
			break
		}
		order.Symbol = instrumentData.Symbol
	}

	return o.Results, nil
}
