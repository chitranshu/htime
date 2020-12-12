# htime
Converts human readable duration to GoLang time.Duration

## Usage

``` sh
go get github.com/chitranshu/htime
```

## Example

``` go
package main

import (
	"log"

	"github.com/chitranshu/htime"
)

func main() {

	log.Println(htime.Duration("30 seconds"))
	log.Println(htime.Duration("30 minutes"))
	log.Println(htime.Duration("3 hours"))
	log.Println(htime.Duration("3 days"))
	log.Println(htime.Duration("3 weeks"))
	log.Println(htime.Duration("3 months"))
	log.Println(htime.Duration("3 years"))

	log.Println(htime.Duration("1 minute and 30 seconds"))
	log.Println(htime.Duration("4 hours, 30 minutes and 10 seconds"))

	// Negative duration - to represent past
	log.Println(htime.Duration(("-1 year")))
	log.Println(htime.Duration(("-10 minutes")))

}

```

#### Output
``` sh
2020/12/11 16:01:52 30s
2020/12/11 16:01:52 30m0s
2020/12/11 16:01:52 3h0m0s
2020/12/11 16:01:52 72h0m0s
2020/12/11 16:01:52 504h0m0s
2020/12/11 16:01:52 2160h0m0s
2020/12/11 16:01:52 26280h0m0s
2020/12/11 16:01:52 1m30s
2020/12/11 16:01:52 4h30m10s
2020/12/11 16:01:52 -8760h0m0s
2020/12/11 16:01:52 -10m0s
```
