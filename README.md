# sumeragi [![GoDoc](https://godoc.org/github.com/1l0/sumeragi?status.svg)](https://godoc.org/github.com/1l0/sumeragi)

##Example
```go
package main

import (
    "fmt"

    "github.com/1l0/sumeragi"
)

func main() {
    now := sumeragi.Now()
    r, k := sumeragi.Era(now)
    fmt.Printf("Now: %v, %s (%s)\n", now, r, k)
}
```
Prints something like:
```
Now: 2675-10-31 08:05:07.701532835 +0900 SUMERAGI, Heisei (平成)
```
