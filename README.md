# Robinhood Client

## About

This client library is based on [`github.com/andrewstuart/go-robinhood`](https://github.com/andrewstuart/go-robinhood/blob/master/README.md). The major differences are:

- Uses OpenAPI generated structs
- Conforms to Robinhood's API, which means most fields will be treated as a `string` type.
- Includes more testing
- Extends pagination

## Notice

If you have used this library before, and use credential caching, you will need
to remove any credential cache and rebuild if you experience errors.

## General usage

```go
o := &robinhood.CredsCacher{
  Creds: &robinhood.OAuth{
    Username: os.Getenv("ROBINHOOD_USERNAME"),
    Password: os.Getenv("ROBINHOOD_PASSWORD"),
  },
}
c, _ := robinhood.Dial(&robinhood.CredsCacher{Creds: o})

i, _ := c.GetInstrumentForSymbol("SPY")
```

## Repository Views

[![HitCount](http://hits.dwyl.com/austin-millan/robinhood-client.svg)](http://hits.dwyl.com/austin-millan/robinhood-client)
