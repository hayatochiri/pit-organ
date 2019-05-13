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

| Implemented             | Method | Path                                                              |
| :---------------------: | :----- | :---------------------------------------------------------------- |
| <ul><li>[x] </li></ul> | GET    | /v3/accounts                                                      |
| <ul><li>[x] </li></ul> | GET    | /v3/accounts/{accountID}                                          |
| <ul><li>[x] </li></ul> | GET    | /v3/accounts/{accountID}/summary                                  |
| <ul><li>[x] </li></ul> | GET    | /v3/accounts/{accountID}/instruments                              |
| <ul><li>[ ] </li></ul> | PATCH  | /v3/accounts/{accountID}/configuration                            |
| <ul><li>[ ] </li></ul> | GET    | /v3/accounts/{accountID}/changes                                  |
| <ul><li>[x] </li></ul> | GET    | /v3/instruments/{instrument}/candles                              |
| <ul><li>[x] </li></ul> | GET    | /v3/instruments/{instrument}/orderBook                            |
| <ul><li>[x] </li></ul> | GET    | /v3/instruments/{instrument}/positionBook                         |
| <ul><li>[x] </li></ul> | POST   | /v3/accounts/{accountID}/orders                                   |
| <ul><li>[ ] </li></ul> | GET    | /v3/accounts/{accountID}/orders                                   |
| <ul><li>[ ] </li></ul> | GET    | /v3/accounts/{accountID}/pendingOrders                            |
| <ul><li>[ ] </li></ul> | GET    | /v3/accounts/{accountID}/orders/{orderSpecifier}                  |
| <ul><li>[ ] </li></ul> | PUT    | /v3/accounts/{accountID}/orders/{orderSpecifier}                  |
| <ul><li>[ ] </li></ul> | PUT    | /v3/accounts/{accountID}/orders/{orderSpecifier}/cancel           |
| <ul><li>[ ] </li></ul> | PUT    | /v3/accounts/{accountID}/orders/{orderSpecifier}/clientExtensions |
| <ul><li>[ ] </li></ul> | GET    | /v3/accounts/{accountID}/positions                                |
| <ul><li>[ ] </li></ul> | GET    | /v3/accounts/{accountID}/openPositions                            |
| <ul><li>[ ] </li></ul> | GET    | /v3/accounts/{accountID}/positions/{instrument}                   |
| <ul><li>[ ] </li></ul> | PUT    | /v3/accounts/{accountID}/positions/{instrument}/close             |
| <ul><li>[x] </li></ul> | GET    | /v3/accounts/{accountID}/pricing                                  |
| <ul><li>[x] </li></ul> | GET    | /v3/accounts/{accountID}/pricing/stream                           |
| <ul><li>[ ] </li></ul> | GET    | /v3/accounts/{accountID}/trades                                   |
| <ul><li>[ ] </li></ul> | GET    | /v3/accounts/{accountID}/openTrades                               |
| <ul><li>[ ] </li></ul> | GET    | /v3/accounts/{accountID}/trades/{tradeSpecifier}                  |
| <ul><li>[ ] </li></ul> | PUT    | /v3/accounts/{accountID}/trades/{tradeSpecifier}/close            |
| <ul><li>[ ] </li></ul> | PUT    | /v3/accounts/{accountID}/trades/{tradeSpecifier}/clientExtensions |
| <ul><li>[ ] </li></ul> | PUT    | /v3/accounts/{accountID}/trades/{tradeSpecifier}/orders           |
| <ul><li>[x] </li></ul> | GET    | /v3/accounts/{accountID}/transactions                             |
| <ul><li>[ ] </li></ul> | GET    | /v3/accounts/{accountID}/transactions/{transactionID}             |
| <ul><li>[x] </li></ul> | GET    | /v3/accounts/{accountID}/transactions/idrange                     |
| <ul><li>[ ] </li></ul> | GET    | /v3/accounts/{accountID}/transactions/sinceid                     |
| <ul><li>[x] </li></ul> | GET    | /v3/accounts/{accountID}/transactions/stream                      |
