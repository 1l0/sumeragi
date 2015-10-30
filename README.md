# sumeragi [![GoDoc](https://godoc.org/github.com/1l0/sumeragi?status.svg)](https://godoc.org/github.com/1l0/sumeragi)

##Example
```go
func main() {
    now := sumeragi.Now()
    r, k := sumeragi.Era(now)
    fmt.Printf("Now: %v, %s (%s)\n", now, r, k)

    utc := time.Date(987, time.January, 1, 0, 0, 0, 0, time.UTC)
    koki := sumeragi.To(utc)
    r, k = sumeragi.Era(utc) // Any time zones are acceptable for Era
    fmt.Printf("UTC: %v\nKoki: %v, %s (%s)\n", utc, koki, r, k)
}
```
Prints something like:
```
Now: 2675-10-31 08:22:19.577885806 +0900 SUMERAGI, Heisei (平成)
UTC: 0987-01-01 00:00:00 +0000 UTC
Koki: 1647-01-01 09:00:00 +0900 SUMERAGI, Eien (永延)
```
