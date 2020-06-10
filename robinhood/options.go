package robinhood

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"
	model "gitlab.com/brokerage-api/robinhood-openapi/openapi"
)

const dateFormat = "2006-01-02"

// Date is a specific json time format for dates only
type Date struct {
	time.Time
}

// NewDate returns a new Date in the local time zone
func NewDate(y, m, d int) Date {
	return Date{time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)}
}

// NewZonedDate returns a date with a zone.
func NewZonedDate(y, m, d int, z *time.Location) Date {
	return Date{time.Date(y, time.Month(m), d, 0, 0, 0, 0, z)}
}

func (d Date) String() string {
	return d.Format(dateFormat)
}

// MarshalJSON implements json.Marshaler
func (d Date) MarshalJSON() ([]byte, error) {
	return []byte("\"" + d.String() + "\""), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (d *Date) UnmarshalJSON(bs []byte) error {
	t, err := time.Parse(dateFormat, strings.Trim(strings.TrimSpace(string(bs)), "\""))
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

// GetOptionChains returns options for the given instruments
func (c *Client) GetOptionChains(is ...*model.InstrumentData) ([]*model.OptionChain, error) {
	s := []string{}
	for _, inst := range is {
		s = append(s, inst.Id)
	}

	var res struct{ Results []*model.OptionChain }

	err := c.GetAndDecode(EPOptions+"chains/?equity_instrument_ids="+strings.Join(s, ","), &res)
	if err != nil {
		return nil, err
	}

	return res.Results, nil
}

type Pager struct {
	Next, Previous string
}

func (p Pager) HasMore() bool {
	return p.Next != ""
}

func (p *Pager) GetNext(c *Client, out interface{}) error {
	if p.Next == "" {
		return io.EOF
	}

	return c.GetAndDecode(p.Next, out)
}

// GetOptionsInstrument returns a list of option-typed instruments given a list of
// expiration dates for a given trade type. The request will continue until the
// provided context is cancelled. This is done to mimic the way the web UI
// fetches many, many options instruments repeatedly, since I haven't yet
// figured out how/when they decide to stop.
func (c *Client) GetOptionsInstrument(ctx context.Context, o model.OptionChain, tradeType string, date Date) ([]*model.OptionInstrument, error) {
	u := fmt.Sprintf(
		"%sinstruments/?chain_id=%s&expiration_dates=%s&state=active&tradability=tradable&type=%s",
		EPOptions,
		o.Id,
		date,
		tradeType,
	)

	rs := []*model.OptionInstrument{}

	var out struct {
		Results []*model.OptionInstrument
		Pager
	}
	err := c.GetAndDecode(u, &out)
	if err != nil {
		return nil, err
	}

	rs = append(rs, out.Results...)

	for out.HasMore() {
		err := out.GetNext(c, &out)
		if err != nil {
			return rs, err
		}
		rs = append(rs, out.Results...)

		select {
		case <-ctx.Done():
			return rs, nil
		default:
		}
	}
	return rs, nil
}

// OIsForDate filters OptionInstruments for expiration date.
func OIsForDate(os []*model.OptionInstrument, d Date) []*model.OptionInstrument {
	out := make([]*model.OptionInstrument, 0, len(os)/6)
	for i := range os {
		parsedDate, err := time.Parse("01/02/06", os[i].ExpirationDate)
		if err != nil {
			parsedDate, err = time.Parse("2006-01-02", os[i].ExpirationDate)
			if err != nil {

			}
		}
		if parsedDate.Equal(d.Time) {
			out = append(out, os[i])
		}
	}
	return out
}

// MarketData returns market data for all the listed Option instruments
func (c *Client) MarketData(opts ...*model.OptionInstrument) ([]*model.MarketData, error) {
	is := make([]string, len(opts))

	for i, o := range opts {
		is[i] = o.Url
	}

	u, err := url.Parse(EPOptionQuote)
	if err != nil {
		return nil, fmt.Errorf("Error: %v", err)
	}

	// Number of instruments to request at once
	num := 30
	// Number of requests this will require to be made
	n := len(is) / num
	if len(is)%num != 0 {
		n++
	}

	rs := []*model.MarketData{}

	for i := 0; i < n; i++ {
		end := (i+1)*num + 1
		if end > len(is) {
			end = len(is)
		}

		q := url.Values{"instruments": []string{strings.Join(is[i*num:end], ",")}}

		u.RawQuery = q.Encode()

		var r struct{ Results []*model.MarketData }
		if e := c.GetAndDecode(u.String(), &r); err != nil {
			err = multierror.Append(err, e)
			continue
		}
		for _, res := range r.Results {
			if res != nil {
				rs = append(rs, res)
			}
		}
	}

	return rs, err
}
