# Golang-exchange-event-listener


Data to subscribe to notifications.  
```{“action”: “subscribe”, “symbols”: <[]string>}```

Attention some quotes do not have a lastPrice field. In this case, you will receive 0 in the price field.

Data to unsubscribe to notifications.  
``` {“action”: “unsubscribe”}```