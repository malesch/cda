package grifts

import (
	"github.com/cdacontrol/cda/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
