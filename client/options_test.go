package robinhood

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMarketData(t *testing.T) {
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

	asrt.NoError(err)
	asrt.NotNil(c)

	i, err := c.GetInstrumentForSymbol("SPY")
	asrt.NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	ch, err := c.GetOptionChains(ctx, i)
	asrt.NoError(err)
	if len(ch) < 1 {
		t.Errorf("Unable to get options chain for instrument")
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	insts, err := c.GetOptionsInstrument(ctx, ch[0], "call", NewDate(2021, 3, 31))
	asrt.NoError(err)

	fmt.Printf("len(insts) = %+v\n", len(insts))

	is, err := c.MarketData(insts...)
	asrt.NoError(err)

	// spew.Dump(is)
	fmt.Printf("len(is) = %+v\n", len(is))
}
