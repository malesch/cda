package grifts

import (
	"github.com/cdacontrol/cda/models"
	"github.com/markbates/grift/grift"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		// Users
		u := &models.User{
			Name:                 "cda",
			Password:             "cda",
			PasswordConfirmation: "cda",
			RoleID:               2, // Admin
			LocaleID:             1, // Deutsch
		}
		return models.DB.Create(u)
	})

})
