package robinhood

import (
	"errors"
	"fmt"

	model "gitlab.com/brokerage-api/robinhood-client/models"
)

// GetCryptoCurrencyPairs will give which crypto currencies are tradeable and corresponding ids
func (c *Client) GetCryptoCurrencyPairs() ([]model.CryptoCurrencyPair, error) {
	var r struct{ Results []model.CryptoCurrencyPair }
	err := c.GetAndDecode(EPCryptoCurrencyPairs, &r)
	return r.Results, err
}

// GetCryptoInstrument will take standard crypto symbol and return usable information
// to place the order
func (c *Client) GetCryptoInstrument(symbol string) (*model.CryptoCurrencyPair, error) {
	allPairs, err := c.GetCryptoCurrencyPairs()
	if err != nil {
		return nil, fmt.Errorf("call failed with error: %v", err.Error())
	}

	for _, pair := range allPairs {
		if pair.AssetCurrency.Code == symbol {
			return &pair, nil
		}
	}

	return nil, errors.New("could not find given symbol")
}
