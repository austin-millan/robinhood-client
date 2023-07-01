package robinhood

import (
	"sync"
	model "gitlab.com/brokerage-api/robinhood-client/models"
)

// GetWatchlists retrieves the watchlists for a given set of credentials/accounts.
func (c *Client) GetWatchlists() ([]model.Watchlist, error) {
	var r struct{ Results []model.Watchlist }
	err := c.GetAndDecode(EPWatchlists, &r)
	if err != nil {
		return nil, err
	}
	return r.Results, nil
}

// GetInstruments returns the list of Instruments associated with a Watchlist.
func (c *Client) GetInstruments(w *model.Watchlist) ([]model.InstrumentData, error) {
	var r struct {
		Results []struct {
			Instrument, URL string
		}
	}

	err := c.GetAndDecode(w.Url, &r)
	if err != nil {
		return nil, err
	}

	insts := make([]*model.InstrumentData, len(r.Results))
	wg := &sync.WaitGroup{}
	wg.Add(len(r.Results))

	for i := range r.Results {
		go func(i int) {
			defer wg.Done()

			inst, err := c.GetInstrument(r.Results[i].Instrument)
			if err != nil {
				return
			}

			insts[i] = inst
		}(i)
	}

	wg.Wait()

	// Filter slice for empties (if error)
	retInsts := []model.InstrumentData{}
	for _, inst := range insts {
		if inst != nil {
			retInsts = append(retInsts, *inst)
		}
	}

	return retInsts, err
}
