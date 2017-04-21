udprouter
---

![](http://i.imgur.com/JgfJxX4.gif)

#### Usage

```go
package main

import (
	"fmt"
	"time"

	udpr "github.com/rikonor/udprouter"
)

func main() {
	r := udpr.NewUDPRouter()
	r.Handle("time_sync", handleTimeSync)
	log.Fatal(r.Listen(":8085"))
}

func handleTimeSync(body string, respond udpr.UDPResponseFunc) {
	t := time.Now().Unix()
	tstr := fmt.Sprintf("%d", t)
	respond([]byte(tstr))
}
```

MIT
