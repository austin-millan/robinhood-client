# Robinhood Client

## About

This client library is based on [`github.com/andrewstuart/go-robinhood`](https://github.com/andrewstuart/go-robinhood/blob/master/README.md). The major differences are:

- Uses OpenAPI generated structs
- Conforms to Robinhood's API, which means most fields will be treated as a `string` type.
- Includes more testing

## Notice

If you have used this library before, and use credential caching, you will need
to remove any credential cache and rebuild if you experience errors.

## General usage

```go
cli, err := robinhood.Dial(&robinhood.OAuth{
  Username: "andrewstuart",
  Password: "mypasswordissecure",
})

//err

i, err := cli.GetInstrumentForSymbol("SPY")

//err

o, err := cli.Order(i, robinhood.OrderOpts{
  Price: 100.0,
  Side: robinhood.Buy,
  Quantity: 1,
})

//err

err := cli.CancelOrder(o)

if err != nil {
  //Oh well
}
```
