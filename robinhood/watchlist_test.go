package robinhood

import (
	"os"
	"testing"
)

func TestGetWatchlists(t *testing.T) {
	if os.Getenv("ROBINHOOD_USERNAME") == "" {
		t.Skip("No username set")
		return
	}
	o := &CredsCacher{
		Creds: &OAuth{
			Username: os.Getenv("ROBINHOOD_USERNAME"),
			Password: os.Getenv("ROBINHOOD_PASSWORD"),
		},
	}
	c, err := Dial(&CredsCacher{Creds: o})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	res, err := c.GetWatchlists()
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	_ = len(res)
}

func TestGetWatchlistInstruments(t *testing.T) {
	if os.Getenv("ROBINHOOD_USERNAME") == "" {
		t.Skip("No username set")
		return
	}
	o := &CredsCacher{
		Creds: &OAuth{
			Username: os.Getenv("ROBINHOOD_USERNAME"),
			Password: os.Getenv("ROBINHOOD_PASSWORD"),
		},
	}
	c, err := Dial(&CredsCacher{Creds: o})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	watchlists, err := c.GetWatchlists()
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if len(watchlists) < 1 {
		t.Errorf("No watchlist to run this test on")
		return
	}
	instrumentWatchlists, err := c.GetInstruments(&watchlists[0])
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	_ = len(instrumentWatchlists)
}
