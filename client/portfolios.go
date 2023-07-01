package robinhood

import (
	"github.com/AlekSi/pointer"
	model "gitlab.com/brokerage-api/robinhood-client/models"
)

// GetPortfolios returns all the portfolios associated with a client's
// credentials and accounts
func (c *Client) GetPortfolios() ([]model.Portfolio, error) {
	var p struct{ Results []model.Portfolio }
	err := c.GetAndDecode(EPPortfolios, &p)
	return p.Results, err
}

// GetCryptoPortfolios returns crypto portfolio info
func (c *Client) GetCryptoPortfolios() (model.CryptoPortfolio, error) {
	var p model.CryptoPortfolio
	var portfolioURL = EPCryptoPortfolio + pointer.GetString(c.CryptoAccount.Id)
	err := c.GetAndDecode(portfolioURL, &p)
	return p, err
}
