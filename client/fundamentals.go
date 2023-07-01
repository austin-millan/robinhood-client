package robinhood

import (
	"strings"

	model "gitlab.com/brokerage-api/robinhood-client/models"
)

// GetFundamentals returns fundemental data for the list of stocks provided.
func (c *Client) GetFundamentals(stocks ...string) ([]model.FundamentalsData, error) {
	url := EPFundamentals + "?symbols=" + strings.Join(stocks, ",")
	var r struct{ Results []model.FundamentalsData }
	err := c.GetAndDecode(url, &r)
	return r.Results, err
}
