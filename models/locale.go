package models

// Locales is the map of available locales
var Locales = map[int]string{
	1: "de",
	2: "fr",
	3: "it",
	4: "en",
}

// SelectLocales returns a map of Locale localization ID's with
// their associated role number
func SelectLocales() map[string]int {
	locs := make(map[string]int)
	for id, name := range Locales {
		locs["locale."+name] = id
	}
	return locs
}
