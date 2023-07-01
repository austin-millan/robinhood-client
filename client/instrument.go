package robinhood

import (
	"fmt"

	model "gitlab.com/brokerage-api/robinhood-client/models"
)

// GetInstrument returns an Instrument given a URL
func (c *Client) GetInstrument(instURL string) (*model.InstrumentData, error) {
	var i model.InstrumentData
	err := c.GetAndDecode(instURL, &i)
	if err != nil {
		return nil, err
	}
	return &i, err
}

// GetInstrumentForSymbol returns an Instrument given a ticker symbol
func (c *Client) GetInstrumentForSymbol(sym string) (*model.InstrumentData, error) {
	var i struct {
		Results []model.InstrumentData
	}
	err := c.GetAndDecode(EPInstruments+"?symbol="+sym, &i)
	if err != nil {
		return nil, err
	}
	if len(i.Results) < 1 {
		return nil, fmt.Errorf("no results")
	}
	return &i.Results[0], err
}
