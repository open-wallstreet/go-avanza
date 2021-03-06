# Websocket

You can connect to Avanza Websocket API using the client. The client will handle connection/pinging.

You can set up a new connection like this:

```go
_, err := client.Auth.Authenticate(context.Background(), username, password, totpSecret)
if err != nil {
    log.Fatalf(err.Error())
}

timeout, _ := context.WithTimeout(context.Background(), 90*time.Second)
_, quotes, err := client.Websocket.StreamQuotes(timeout, "19000") // 19000 = USD/SEK
if err != nil {
    log.Fatalf(err.Error())
}
for q := range quotes {
    log.Println(q)
}
```

Websocket stream methods will return a channel containing all messages that will come in. 
If the context is closed we will close the channel and continue your program.

### Subscription strings

Here is a quick explanation how to build the subscription params string.



|                    | Single          | Multiple                        |
|--------------------|:----------------|:--------------------------------|
| Quotes             | `{orderbookId}` | `{orderbookId1},{orderbookId2}` |
| OrderDepth         | `{orderbookId}` | `{orderbookId1},{orderbookId2}` |
| BrokerTradeSummary | `{orderbookId}` | `{orderbookId1},{orderbookId2}` |
| Trades             | `{orderbookId}` | `{orderbookId},{orderbookId}`   |
| Positions          | `{accountID}`   | `{accountID1},{accountID2}`     |
| Orders             | `{orderID}`     | `{orderID1},{orderID2}`         |