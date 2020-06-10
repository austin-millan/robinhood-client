package robinhood

import (
	"context"
	"os"
	"testing"
	"time"

	// "time"

	// "github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	model "gitlab.com/brokerage-api/robinhood-openapi/openapi"
)

// GetOptionsOrders returns all outstanding options orders
func TestGetOptionsOrders(t *testing.T) {
	if os.Getenv("ROBINHOOD_USERNAME") == "" {
		t.Skip("No username set")
		return
	}
	// asrt := assert.New(t)
	o := &OAuth{
		Username: os.Getenv("ROBINHOOD_USERNAME"),
		Password: os.Getenv("ROBINHOOD_PASSWORD"),
	}

	c, err := Dial(&CredsCacher{Creds: o})

	res, err := c.GetOptionsOrders()
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	_ = res
}

// TODO
func TestOrderOptions(t *testing.T) {
	if os.Getenv("ROBINHOOD_USERNAME") == "" {
		t.Skip("No username set")
		return
	}
	asrt := assert.New(t)
	o := &OAuth{
		Username: os.Getenv("ROBINHOOD_USERNAME"),
		Password: os.Getenv("ROBINHOOD_PASSWORD"),
	}

	c, err := Dial(&CredsCacher{Creds: o})
	i, err := c.GetInstrumentForSymbol("SPY")
	asrt.NoError(err)

	ch, err := c.GetOptionChains(i)
	asrt.NoError(err)
	if len(ch) < 1 {
		t.Errorf("Unable to get options chain for instrument")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	insts, err := c.GetOptionsInstrument(ctx, *ch[0], "call", NewDate(2021, 3, 31))
	asrt.NoError(err)
	inputOpts := OptionsOrderOpts{
		Side:        model.BUY,
		Type:        model.MARKET,
		Quantity:    1,
		Direction:   model.DEBIT,
		Price:       0.01,
		TimeInForce: model.GTC,
	}
	res, err := c.OrderOptions(insts[0], inputOpts)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	_ = res
}
