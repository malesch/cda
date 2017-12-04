package models

import (
	"log"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/markbates/pop"
)

// DB is a connection to your database to be used
// throughout your application.
var DB *pop.Connection

func init() {
	var err error
	env := envy.Get("GO_ENV", "development")
	DB, err = pop.Connect(env)
	if err != nil {
		log.Fatal(err)
	}
	pop.Debug = env == "development"
}

func InjectModelGlobals(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {

		c.Set("selectLocales", SelectLocales(c))
		c.Set("selectRoles", SelectRoles(c))

		c.Logger().Info(c.Session())
		c.Logger().Info(SelectRoles(c))
		c.Logger().Info("*********----> " + Translate(c, "welcome_greeting"))

		return next(c)
	}
}

func Translate(ctx buffalo.Context, key string) string {
	tfn, ok := ctx.Value("t").(func(string) (string, error))
	if !ok {
		return key
	}
	msg, err := tfn(key)
	if err != nil {
		return key
	}

	return msg
}
