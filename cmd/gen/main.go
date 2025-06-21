// generator of barebone eras.go
// a json file is retrieved from https://gengoh.jp/index.html

package main

import (
	"cmp"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"
)

type body struct {
	Gengoh map[string]any `json:"gengoh"`
}

type gengou struct {
	ID       float64 `json:"id"`
	Kanji    string  `json:"name"`
	Hiragana string  `json:"yomi"`
	Romaji   string  `json:"rome"`
	Epoch    string  `json:"bgn"`
}

const header = `package sumeragi

import "time"

var eras = []Era{
	{
		Kanji:    "",
		Romaji:   "",
		Hiragana: "",
		Epoch:    time.Date(-660, 1, 1, 0, 0, 0, 0, JST),
		Emperors: []string{""},
	},
	`

const footer = `
}`

const tmpl = `{
		Kanji:    "%s",
		Hiragana: "%s",
		Romaji:   "%s",
		Epoch:    time.Date(%s, %s, %s, 0, 0, 0, 0, JST),
		Emperors: []string{""},
	},
	`

func main() {
	f, err := os.Open("testdata/gengou4.json")
	if err != nil {
		panic(err)
	}
	b := body{}
	if err := json.NewDecoder(f).Decode(&b); err != nil {
		panic(err)
	}
	if err := f.Close(); err != nil {
		panic(err)
	}

	gengoh := b.Gengoh

	var gengous []gengou

	for _, v := range gengoh {
		value := make(map[string]any, 0)
		if vv, ok := v.(map[string]any); ok {
			value = vv
		}
		id := value["id"].(float64)
		name := value["name"].(string)
		yomi := value["yomi"].(string)
		rome := value["rome"].(string)
		bgn := value["bgn"].(string)
		gengous = append(gengous, gengou{
			ID:       id,
			Kanji:    name,
			Hiragana: yomi,
			Romaji:   rome,
			Epoch:    bgn,
		})
	}

	slices.SortFunc(gengous, func(a, b gengou) int {
		return cmp.Compare(a.ID, b.ID)
	})

	out, err := os.Create("eras.go")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	output := header
	for _, g := range gengous {
		ts := strings.Split(g.Epoch, "-")
		if len(ts) != 3 {
			panic(fmt.Errorf("invalid bgn format"))
		}
		year := strings.TrimLeft(ts[0], "0")
		month := strings.TrimLeft(ts[1], "0")
		date := strings.TrimLeft(ts[2], "0")
		output += fmt.Sprintf(tmpl, g.Kanji, g.Hiragana, g.Romaji, year, month, date)
	}
	output += footer
	if _, err := out.WriteString(output); err != nil {
		panic(err)
	}
}
