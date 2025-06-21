package sumeragi

import (
	"fmt"
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	now := Now()
	zone, _ := now.Zone()
	if zone != "SUMERAGI" {
		t.Error("TIme zone must be SUMERAGI")
	}
}

func TestEra(t *testing.T) {
	m := time.Date(512, time.January, 1, 0, 0, 0, 0, time.UTC)
	e := ToEra(m)
	if e.Kanji != "" {
		t.Error("era must be unnamed")
	}
	m = time.Date(1597, time.January, 1, 0, 0, 0, 0, time.UTC)
	e = ToEra(m)
	if e.Kanji != "慶長" {
		t.Error("invalid era")
	}
	m = time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	e = ToEra(m)
	if e.Kanji != "令和" {
		t.Error("invalid era")
	}
}

func TestEraFromName(t *testing.T) {
	name := "元禄"
	e := EraFromName(name)
	if e.Epoch.Year() != 1688 {
		t.Error("mismatch era")
	}
	name = "れいわ"
	e = EraFromName(name)
	if e.Epoch.Year() != 2019 {
		t.Error("mismatch era")
	}
}

func Example() {
	now := Now()
	e := ToEra(now)
	fmt.Printf("Now: (Sumeragi) %v, Era: %v\n", now, e)
}
