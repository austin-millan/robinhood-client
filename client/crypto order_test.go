package robinhood

import (
	"os"

	"testing"

	model "gitlab.com/brokerage-api/robinhood-client/models"
)

func TestCryptoOrder(t *testing.T) {
	if os.Getenv("ROBINHOOD_USERNAME") == "" {
		t.Skip("No username set")
	}
	o := &CredsCacher{
		Creds: &OAuth{
			Username: os.Getenv("ROBINHOOD_USERNAME"),
			Password: os.Getenv("ROBINHOOD_PASSWORD"),
		},
	}
	c, err := Dial(&CredsCacher{Creds: o})
	res, err := c.GetCryptoInstrument("BTC")
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if res == nil {
		t.Errorf("Couldn't get symbol")
		return
	}
	inputOpts := CryptoOrderOpts{
		Side:            model.SELL,
		Type:            model.MARKET,
		Quantity:        1,
		AmountInDollars: 5,
		Price:           40000.00,
		TimeInForce:     model.GTC,
	}
	output, err := c.CryptoOrder(*res, inputOpts)
	if err != nil {
		t.Errorf("Skipping error: %v", err)
	}
	if output == nil {
		t.Errorf("No output")
	}
}
