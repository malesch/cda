package actions

import (
	"github.com/gobuffalo/buffalo"
)

func LocalizeSelect(c buffalo.Context, selectMap map[string]int) map[string]int {
	locSelectMap := make(map[string]int)
	for name, id := range selectMap {
		locSelectMap[T.Translate(c, name)] = id
	}
	return locSelectMap
}
