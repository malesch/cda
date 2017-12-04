package models

import "github.com/gobuffalo/buffalo"

type Locale struct {
	ID   int
	Name string
}

var Locales = []Locale{
	Locale{ID: 0, Name: "de"},
	Locale{ID: 1, Name: "fr"},
	Locale{ID: 2, Name: "it"},
	Locale{ID: 3, Name: "en"},
}

func SelectLocales(c buffalo.Context) map[string]int {
	locs := make(map[string]int)
	for _, loc := range Locales {
		localizedName := Translate(c, "locale."+loc.Name)
		locs[localizedName] = loc.ID
	}
	return locs
}

// func LocaleName(id int) string {
// 	for _, loc := range Locales {
// 		if loc.ID == id {
// 			return loc.Name
// 		}
// 	}
// 	return ""
// }
