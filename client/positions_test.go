package robinhood

import (
	"os"
	"testing"
)

func TestGetPositions(t *testing.T) {
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
	accs, err := c.GetAccounts()
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if len(accs) < 1 {
		t.Errorf("Unable to get accounts")
		return
	}
	res, err := c.GetPositions(accs[0])
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	_ = len(res)
}

