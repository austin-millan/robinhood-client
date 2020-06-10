package main

import (
	"fmt"
	"os"

	"gitlab.com/brokerage-api/robinhood-client/robinhood"
)

func main() {
	if os.Getenv("ROBINHOOD_USERNAME") == "" {
		return
	}
	o := &robinhood.CredsCacher{
		Creds: &robinhood.OAuth{
			Username: os.Getenv("ROBINHOOD_USERNAME"),
			Password: os.Getenv("ROBINHOOD_PASSWORD"),
		},
	}
	c, err := robinhood.Dial(&robinhood.CredsCacher{Creds: o})

	//err
	i, err := c.GetInstrumentForSymbol("SPY")
	if err != nil {
		fmt.Errorf("Unable to gget SPY instrument")
		return
	}
	fmt.Printf("%v", i)
}
