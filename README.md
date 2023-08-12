# blastr

Go module for sending nostr events to [blastr](https://github.com/MutinyWallet/blastr)

```go
package main

import (
	"context"

	"github.com/stemstr/blastr"
)

func main() {
	client, _ := blastr.New("nsec1xxx")

	_ = client.SendText(context.Background(), "hello, world!")
}
```
