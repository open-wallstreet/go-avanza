# Avanza Unofficial GO API Client
[![Go Reference](https://pkg.go.dev/badge/github.com/open-wallstreet/go-avanza.svg)](https://pkg.go.dev/github.com/open-wallstreet/go-avanza)
![CI](https://github.com/open-wallstreet/go-avanza/actions/workflows/build_and_test.yml/badge.svg)

**Please note that I am not affiliated with Avanza Bank AB in any way. The underlying API can be taken down or changed without warning at any point in time.**

## Installation

`go get github.com/open-wallstreet/go-avanza`

## Getting a TOTP Secret

**NOTE: Since May 2018 two-factor authentication is used to log in.**

Here are the steps to get your TOTP Secret:

0. Go to Mina Sidor > Profil > Sajtinställningar > Tvåfaktorsinloggning and click "Återaktivera". (_Only do this step if you have already set up two-factor auth._)
1. Click "Aktivera" on the next screen.
2. Select "Annan app för tvåfaktorsinloggning".
3. Click "Kan du inte scanna QR-koden?" to reveal your TOTP Secret.
4. That codes should be passed into the `Authenticate` method. Make sure not to save it into your codebase. Instead supply it trough environment variables or other encoded systems.

## Documentation

Docs: [![Go Reference](https://pkg.go.dev/badge/github.com/open-wallstreet/go-avanza.svg)](https://pkg.go.dev/github.com/open-wallstreet/go-avanza)

[Websocket information](./docs/websocket.md)

## Quick Guide

```go
// You can create a new client simply like this
func main() {
    client := goavanza.New()
    defer client.Close()
}
```

```go
// Or if you need to debug http responses
func main() {
    client := goavanza.New(goavanza.WithDebug(true))
    defer client.Close()
}
```

Most API calls need you to authenticate using a TOPT token. see [Getting a TOTP Secret](#Getting a TOTP Secret) section how to create one.
After that call the `Authenticate method`
```go
authenticate, err := client.Auth.Authenticate(context.Background(), username, password, totpSecret)
```

**You should not save your username, password or totpSecret in your code. You ENV variables or other ways to encrypt or hide the data**

See `/examples` or [GoDocs](https://pkg.go.dev/github.com/open-wallstreet/go-avanza) for more information

## CLI tool
You can install the CLI tool by running

```bash
go install github.com/open-wallstreet/go-avanza
```

Afterwards you can download complete Avanza stock list by running the download command

```bash
go-avanza downloader stocks-list -o myfile.csv
```


### Shoutouts

Major inspiration goes to [Node.js unofficial avanza api](https://github.com/fhqvst/avanza)

### RESPONSIBILITIES

The author of this software is not responsible for any indirect damages (foreseeable or unforeseeable), such as, if necessary, loss or alteration of or fraudulent access to data, accidental transmission of viruses or of any other harmful element, loss of profits or opportunities, the cost of replacement goods and services or the attitude and behavior of a third party.


### Known issues

- GetMarketData response does not have correct interface mapping due to not being able to seen a response with data in it yet.
- Most streaming data has not been mapped to correct structs due to not being able to test responses due to Avanza being down for maintenance during development
