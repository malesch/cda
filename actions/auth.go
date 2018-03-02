package actions

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/cdacontrol/cda/models"
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// AuthNew load the signIn page if not logged in or forward to index
func AuthNew(c buffalo.Context) error {
	c.Set("user", models.User{})
	return c.Render(200, r.HTML("login.html"))
}

// AuthCreate attempts to log the user in with an existing account.
func AuthCreate(c buffalo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return errors.WithStack(err)
	}

	tx := c.Value("tx").(*pop.Connection)

	// find a user with the email
	err := tx.Where("name = ?", strings.ToLower(u.Name)).First(u)

	// helper function to handle bad attempts
	bad := func() error {
		c.Set("user", u)
		verrs := validate.NewErrors()
		verrs.Add("name", "Falscher Benutzername oder Passwort!")
		c.Set("errors", verrs)
		return c.Render(422, r.HTML("login.html"))
	}

	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			// couldn't find an user with the supplied email address.
			return bad()
		}
		return errors.WithStack(err)
	}

	// confirm that the given password matches the hashed password from the db
	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(u.Password))
	if err != nil {
		return bad()
	}

	c.Session().Set("current_user_id", u.ID)
	c.Flash().Add("success", fmt.Sprintf("Welcome, %s!", u.Name))

	return c.Redirect(302, "/")
}

// AuthDestroy clears the session and logs a user out
func AuthDestroy(c buffalo.Context) error {
	c.Session().Clear()
	c.Flash().Add("success", "You successfully logged out!")
	return c.Redirect(302, "/")
}