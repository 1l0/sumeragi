# sumeragi [![GoDoc](https://godoc.org/github.com/1l0/sumeragi?status.svg)](https://godoc.org/github.com/1l0/sumeragi)

Get 皇暦 or 元号.

## Example
```go
d := time.Date(1704, time.May, 1, 0, 0, 0, 0, time.UTC)
e := sumeragi.EraFromTime(d)
sumeragiDate := sumeragi.To(e.Epoch)
fmt.Printf("%s %s %s\n", e.Kanji, e.Emperors[0], sumeragiDate)

// Output: 宝永 中御門天皇 2364-04-16 00:00:00 +0900 SUMERAGI
```
