package robinhood

import (
	"bytes"
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"encoding/json"
	"net/http"

	model "gitlab.com/brokerage-api/robinhood-client/models"
)

// CryptoOrderOpts encapsulates differences between order types
type CryptoOrderOpts struct {
	Side            model.Side
	Type            model.ExecutionType
	AmountInDollars float64
	Quantity        float64
	Price           float64
	TimeInForce     model.TimeInForce
	ExtendedHours   bool
	Stop, Force     bool
}

// CryptoOrder will actually place the order
func (c *Client) CryptoOrder(cryptoPair model.CryptoCurrencyPair, o CryptoOrderOpts) (*model.CryptoOrderOutput, error) {
	var quantity = 0.00
	if o.Price > 0 {
		quantity = o.AmountInDollars / o.Price
	}
	a := model.CryptoOrder{
		AccountId:      c.CryptoAccount.Id,
		CurrencyPairId: cryptoPair.Id,
		Quantity:       fmt.Sprintf("%f", quantity),
		Price:          fmt.Sprintf("%f", o.Price),
		RefId:          uuid.New().String(),
		Side:           o.Side,
		TimeInForce:    o.TimeInForce,
		Type:           o.Type,
	}

	payload, err := json.Marshal(a)

	if err != nil {
		return nil, err
	}

	post, err := http.NewRequest("POST", EPCryptoOrders, bytes.NewReader(payload))
	post.Header.Add("Content-Type", "application/json")

	var out model.CryptoOrderOutput
	err = c.DoAndDecode(post, &out)

	if err != nil {
		return nil, err
	}

	return &out, nil
}

// CancelCryptoOrder to cancel order
func (c *Client) CancelCryptoOrder(o model.CryptoOrderOutput) error {
	post, err := http.NewRequest("POST", o.CancelUrl, nil)
	if err != nil {
		return err
	}

	var output model.CryptoOrderOutput
	err = c.DoAndDecode(post, &output)

	if err != nil {
		return errors.Wrap(err, "could not decode response")
	}

	if output.RejectReason != "" {
		return errors.New(output.RejectReason)
	}

	return nil
}
