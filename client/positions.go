package robinhood

import (
	"net/url"

	"github.com/AlekSi/pointer"
	model "gitlab.com/brokerage-api/robinhood-client/models"
)

// GetPositions returns all the positions associated with an account.
func (c *Client) GetPositions(a model.AccountInfo) ([]model.Position, error) {
	output, err := c.GetPositionsParams(a, PositionParams{})
	return output, err
}

// PositionParams encapsulates parameters known to the RobinHood positions API
// endpoint.
type PositionParams struct {
	NonZero bool
}

// Encode returns the query string associated with the requested parameters
func (p PositionParams) encode() string {
	v := url.Values{}
	if p.NonZero {
		v.Set("nonzero", "true")
	}
	return v.Encode()
}

// GetPositionsParams returns all the positions associated with a count, but
// passes the encoded PositionsParams object along to the RobinHood API as part
// of the query string.
func (c *Client) GetPositionsParams(a model.AccountInfo, p PositionParams) ([]model.Position, error) {
	u, err := url.Parse(pointer.GetString(a.Positions))
	if err != nil {
		return nil, err
	}
	u.RawQuery = p.encode()

	var r struct{ Results []model.Position }
	return r.Results, c.GetAndDecode(u.String(), &r)
}
