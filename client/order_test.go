package robinhood

import (
	"os"
	"testing"
)

func TestRecentOrders(t *testing.T) {
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
	res, err := c.RecentOrders()
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	_ = len(res)
}

func TestGetStockOrders(t *testing.T) {
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
	res, err := c.GetStockOrders()
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	_ = len(res)
}