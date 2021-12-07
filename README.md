# Arriba Control

Arriba Control is a software for manage the transactions of Buy/Sell and Transfer of crypto like BTC and ETH from USD Fiat.

## Demo Online

The app is deployed on Google App Engine

`URL`: https://sage-surfer-260.appspot.com/

## Installation

First we need install all dependencies
```shell
$ go get tidy
```

Now we can run the server with:
```shell
$ go run ./api/main.go
```
The app will be available on `http://localhost:8080`

## Tests

You can run the test of all directories with:
```shell
$ go test ./...
```

## Endpoints

You can find the insomnia file [here](./arriba_endpoints.json) import this file on Postman or Insomnia for handle every endpoint.

>**Note**: All **POST** endpoint must have a header called `idempotent_key` for avoid inconsistent states. 