// Package sumeragi implements time in Japanese calendar(皇暦, 皇紀, 元号).
package sumeragi

import (
	"time"
)

var (
	SUMERAGI = time.FixedZone("SUMERAGI", 32400)
	JST      = time.FixedZone("JST", 32400)
)

// Era represents Japanese era.
type Era struct {
	Kanji    string
	Hiragana string
	Romaji   string
	Epoch    time.Time
	Emperors []string
}

// Now in Japanese calendar.
func Now() time.Time {
	return time.Now().In(SUMERAGI).AddDate(660, 0, 0)
}

// To time in Japanese calendar.
func To(t time.Time) time.Time {
	if z, _ := t.Zone(); z == "SUMERAGI" {
		return t
	}
	return t.In(SUMERAGI).AddDate(660, 0, 0)
}

// ToEra returns an era of the time.
func ToEra(t time.Time) Era {
	for i, era := range eras {
		if i == 0 {
			continue
		}
		if era.Epoch.After(t) {
			return eras[i-1]
		}
	}
	return eras[len(eras)-1]
}

func EraFromName(name string) *Era {
	for i, era := range eras {
		if i == 0 {
			continue
		}
		if era.Kanji == name || era.Hiragana == name {
			return &era
		}
	}
	return nil
}
