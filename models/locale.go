package models

var Locales = map[int]string{
	0: "de",
	1: "fr",
	2: "it",
	3: "en",
}

func SelectLocales() map[string]int {
	locs := make(map[string]int)
	for id, name := range Locales {
		locs["locale."+name] = id
	}
	return locs
}
