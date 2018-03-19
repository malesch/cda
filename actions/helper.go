package actions

import "github.com/gobuffalo/buffalo"

// LocalizeSelect returns a new version of the given map with localized keys
func LocalizeSelect(c buffalo.Context, selectMap map[string]int) map[string]int {
	locSelectMap := make(map[string]int)
	for name, id := range selectMap {
		locName := T.Translate(c, name)
		locSelectMap[locName] = id
	}
	return locSelectMap
}
