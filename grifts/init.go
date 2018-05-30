package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/malesch/cda/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
