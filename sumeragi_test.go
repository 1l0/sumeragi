package sumeragi

import (
	"fmt"
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	now := Now()
	zone, _ := now.Zone()
	t.Log(now)
	if zone != "SUMERAGI" {
		t.Error("TIme zone must be SUMERAGI")
	}
}

func TestEra(t *testing.T) {
	m := time.Date(987, time.January, 1, 0, 0, 0, 0, time.UTC)
	r, k := Era(m)
	t.Log(m, "=", r, k)
	if r != "Eien" {
		t.Error("Era must be Eien")
	}
}

func Example() {
	now := Now()
	r, _ := Era(now)
	fmt.Printf("Now: Sumeragi(Koki) %v, %s Era\n", now, r)
}
