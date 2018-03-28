package actions

import "github.com/gobuffalo/buffalo"

// LocalizeSelect return a map for localizing selection values
func LocalizeSelect(c buffalo.Context, selectMap map[string]int) map[string]int {
	locSelectMap := make(map[string]int)
	for name, id := range selectMap {
		locName := T.Translate(c, name)
		locSelectMap[locName] = id
	}
	return locSelectMap
}
