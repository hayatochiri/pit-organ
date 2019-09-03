# Pit-organ

Golang OANDA v20 REST API wrapper.

## Development

```sh
$ cd /path/to/here
$ docker-compose up --build -d

$ ./run dep ensure
$ ./run gotest
$ ./run go build
$ ./run [command]
```

## API

| Implemented             | Method | Path                                                                                                                                |
| :---------------------: | :----- | :---------------------------------------------------------------------------------------------------------------------------------- |
| <ul><li>[x] </li></ul>  | GET    | [/v3/accounts](https://godoc.org/github.com/hayatochiri/pit-organ#ReceiverAccounts.Get)                                             |
| <ul><li>[x] </li></ul>  | GET    | [/v3/accounts/{accountID}](https://godoc.org/github.com/hayatochiri/pit-organ#ReceiverAccountID.Get)                                |
| <ul><li>[x] </li></ul>  | GET    | [/v3/accounts/{accountID}/summary](https://godoc.org/github.com/hayatochiri/pit-organ#ReceiverAccountSummary.Get)                   |
| <ul><li>[x] </li></ul>  | GET    | [/v3/accounts/{accountID}/instruments](https://godoc.org/github.com/hayatochiri/pit-organ#ReceiverAccountInstruments.Get)           |
| <ul><li>[x] </li></ul>  | PATCH  | [/v3/accounts/{accountID}/configuration](https://godoc.org/github.com/hayatochiri/pit-organ#ReceiverAccountConfiguration.Patch)     |
| <ul><li>[x] </li></ul>  | GET    | [/v3/accounts/{accountID}/changes](https://godoc.org/github.com/hayatochiri/pit-organ#ReceiverAccountChanges.Get)                   |
| <ul><li>[x] </li></ul>  | GET    | [/v3/instruments/{instrument}/candles](https://godoc.org/github.com/hayatochiri/pit-organ#ReceiverInstrumentCandles.Get)            |
| <ul><li>[x] </li></ul>  | GET    | [/v3/instruments/{instrument}/orderBook](https://godoc.org/github.com/hayatochiri/pit-organ#ReceiverInstrumentOrderBook.Get)        |
| <ul><li>[x] </li></ul>  | GET    | [/v3/instruments/{instrument}/positionBook](https://godoc.org/github.com/hayatochiri/pit-organ#ReceiverInstrumentPositionBook.Get)  |
| <ul><li>[x] </li></ul>  | POST   | [/v3/accounts/{accountID}/orders](https://godoc.org/github.com/hayatochiri/pit-organ#ReceiverOrders.Post)                           |
| <ul><li>[x] </li></ul>  | GET    | [/v3/accounts/{accountID}/orders](https://godoc.org/github.com/hayatochiri/pit-organ#ReceiverOrders.Get)                           |
| <ul><li>[ ] </li></ul>  | GET    | /v3/accounts/{accountID}/pendingOrders                                                                                              |
| <ul><li>[ ] </li></ul>  | GET    | /v3/accounts/{accountID}/orders/{orderSpecifier}                                                                                    |
| <ul><li>[ ] </li></ul>  | PUT    | /v3/accounts/{accountID}/orders/{orderSpecifier}                                                                                    |
| <ul><li>[ ] </li></ul>  | PUT    | /v3/accounts/{accountID}/orders/{orderSpecifier}/cancel                                                                             |
| <ul><li>[ ] </li></ul>  | PUT    | /v3/accounts/{accountID}/orders/{orderSpecifier}/clientExtensions                                                                   |
| <ul><li>[ ] </li></ul>  | GET    | /v3/accounts/{accountID}/positions                                                                                                  |
| <ul><li>[ ] </li></ul>  | GET    | /v3/accounts/{accountID}/openPositions                                                                                              |
| <ul><li>[ ] </li></ul>  | GET    | /v3/accounts/{accountID}/positions/{instrument}                                                                                     |
| <ul><li>[ ] </li></ul>  | PUT    | /v3/accounts/{accountID}/positions/{instrument}/close                                                                               |
| <ul><li>[x] </li></ul>  | GET    | [/v3/accounts/{accountID}/pricing](https://godoc.org/github.com/hayatochiri/pit-organ#ReceiverPricing.Get)                          |
| <ul><li>[x] </li></ul>  | GET    | [/v3/accounts/{accountID}/pricing/stream](https://godoc.org/github.com/hayatochiri/pit-organ#ReceiverPricingStream.Get)             |
| <ul><li>[ ] </li></ul>  | GET    | /v3/accounts/{accountID}/trades                                                                                                     |
| <ul><li>[ ] </li></ul>  | GET    | /v3/accounts/{accountID}/openTrades                                                                                                 |
| <ul><li>[ ] </li></ul>  | GET    | /v3/accounts/{accountID}/trades/{tradeSpecifier}                                                                                    |
| <ul><li>[ ] </li></ul>  | PUT    | /v3/accounts/{accountID}/trades/{tradeSpecifier}/close                                                                              |
| <ul><li>[ ] </li></ul>  | PUT    | /v3/accounts/{accountID}/trades/{tradeSpecifier}/clientExtensions                                                                   |
| <ul><li>[ ] </li></ul>  | PUT    | /v3/accounts/{accountID}/trades/{tradeSpecifier}/orders                                                                             |
| <ul><li>[x] </li></ul>  | GET    | [/v3/accounts/{accountID}/transactions](https://godoc.org/github.com/hayatochiri/pit-organ#ReceiverTransactions.Get)                |
| <ul><li>[ ] </li></ul>  | GET    | /v3/accounts/{accountID}/transactions/{transactionID}                                                                               |
| <ul><li>[x] </li></ul>  | GET    | [/v3/accounts/{accountID}/transactions/idrange](https://godoc.org/github.com/hayatochiri/pit-organ#ReceiverTransactionsIdrange.Get) |
| <ul><li>[ ] </li></ul>  | GET    | /v3/accounts/{accountID}/transactions/sinceid                                                                                       |
| <ul><li>[x] </li></ul>  | GET    | [/v3/accounts/{accountID}/transactions/stream](https://godoc.org/github.com/hayatochiri/pit-organ#ReceiverTransactionsStream.Get)   |
