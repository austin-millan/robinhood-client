package robinhood

import (
	"strings"
	model "gitlab.com/brokerage-api/robinhood-openapi/openapi"
)

// GetQuote returns all the latest stock quotes for the list of stocks provided
func (c *Client) GetQuote(stocks ...string) ([]model.Quote, error) {
	url := EPQuotes + "?symbols=" + strings.Join(stocks, ",")
	var r struct{ Results []model.Quote }
	err := c.GetAndDecode(url, &r)
	return r.Results, err
}
