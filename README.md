# Avanza Unofficial GO API Client

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
4. Finally, run `go run github.com/open-wallstreet/go-avanza/cmd/totp 'PASTE_YOUR_TOTP_SECRET_HERE'` to generate an initial code.
5. Done! From now on all you have to do is supply your secret in with `AVANZA_TOTP_SECRET` environment variable function as in the example below.

## Roadmap

- [x] Authenticate / Reauthenticate
- [x] GetPositions
- [x] GetOverview
- [x] GetAccountOverview
- [x] GetDealsAndOrders
- [x] GetTransactions
- [ ] GetWatchlists
- [ ] AddToWatchlist
- [ ] GetInstrument
- [ ] GetOrderbook
- [ ] GetOrderbooks
- [ ] GetInspirationLists
- [ ] GetInspirationList
- [x] GetOrder
- [x] EditOrder
- [x] DeleteOrder
- [ ] Websocket RealTime data
  - [x] Orders
  - [ ] Quotes

### Shoutouts

Major inspiration goes to [Node.js unofficial avanza api](https://github.com/fhqvst/avanza)

### RESPONSIBILITIES

The author of this software is not responsible for any indirect damages (foreseeable or unforeseeable), such as, if necessary, loss or alteration of or fraudulent access to data, accidental transmission of viruses or of any other harmful element, loss of profits or opportunities, the cost of replacement goods and services or the attitude and behavior of a third party.
