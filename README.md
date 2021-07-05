# DMS - Dead Mans Switch (kinda)

## Usage 

```
go get -u github.com/Schmenn/dms
```

### Example - Gin

```go
package main

import (
	"os"
	"time"

	"github.com/Schmenn/dms"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func main() {
	// new gin router
	r := gin.Default()

	// make dead mans switch
	d := dms.DeadManSwitch{
		TimerDuration: time.Second * 5,
		Username:      "coolUsername",
		Password:      "password123",
		OnTrigger: func() {
			color.Cyan(time.Now().Format("[DMS: 15:04:05.000]") + " Deploy the cows")
			os.Exit(0)
		},
	}
    // register endpoint with the dead mans switch
	r.GET("/dms", gin.WrapF(d.Handler()))

	r.Run(":80")
}
```

## Why?

I was bored 



