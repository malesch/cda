package actions

import (
	"database/sql"

	"github.com/cdacontrol/cda/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
)

// UsersResource is the resource for the User model
type UsersResource struct {
	buffalo.Resource
}

// List gets all Users. This function is mapped to the path
// GET /users
func (v UsersResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	users := &models.Users{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Users from the DB
	if err := q.All(users); err != nil {
		return errors.WithStack(err)
	}

	// Make Users available inside the html template
	c.Set("users", users)

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)
	prepareSelectLists(c)

	return c.Render(200, r.HTML("users/index.html"))
}

// Show gets the data for one User. This function is mapped to
// the path GET /users/{user_id}
func (v UsersResource) Show(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	user := &models.User{}

	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(404, err)
	}

	c.Set("user", user)
	prepareSelectLists(c)

	return c.Render(200, r.HTML("users/show.html"))
}

// New renders the form for creating a new User.
// This function is mapped to the path GET /users/new
func (v UsersResource) New(c buffalo.Context) error {
	c.Set("user", &models.User{})
	prepareSelectLists(c)

	return c.Render(200, r.HTML("users/new.html"))
}

// Create adds a User to the DB. This function is mapped to the path POST /users
func (v UsersResource) Create(c buffalo.Context) error {
	user := &models.User{}

	if err := c.Bind(user); err != nil {
		return errors.WithStack(err)
	}

	tx := c.Value("tx").(*pop.Connection)

	verrs, err := tx.ValidateAndCreate(user)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		c.Set("user", user)
		c.Set("errors", verrs)

		// Render again the new.html template that the user can correct the input.
		return c.Render(422, r.HTML("users/new.html"))
	}

	c.Flash().Add("success", "User was created successfully")

	return c.Redirect(302, "/users/%s", user.ID)
}

// Edit renders a edit form for a User. This function is mapped to the path GET /users/{user_id}/edit
func (v UsersResource) Edit(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	user := &models.User{}

	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(404, err)
	}

	c.Set("user", user)
	prepareSelectLists(c)

	return c.Render(200, r.HTML("users/edit.html"))
}

// Update changes a User in the DB. This function is mapped to the path PUT /users/{user_id}
func (v UsersResource) Update(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	user := &models.User{}

	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(404, err)
	}

	if err := c.Bind(user); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(user)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		c.Set("user", user)
		c.Set("errors", verrs)

		return c.Render(422, r.HTML("users/edit.html"))
	}

	c.Flash().Add("success", "User was updated successfully")

	return c.Redirect(302, "/users/%s", user.ID)
}

// Destroy deletes a User from the DB. This function is mapped to the path DELETE /users/{user_id}
func (v UsersResource) Destroy(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	user := &models.User{}

	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(user); err != nil {
		return errors.WithStack(err)
	}

	c.Flash().Add("success", "User was destroyed successfully")

	return c.Redirect(302, "/users")
}

func prepareSelectLists(c buffalo.Context) {
	c.Set("locales", models.Locales)
	c.Set("selectLocales", LocalizeSelect(c, models.SelectLocales()))
	c.Set("roles", models.Roles)
	c.Set("selectRoles", LocalizeSelect(c, models.SelectRoles()))
}

// SetCurrentUser attempts to find a user based on the current_user_id
// in the session. If one is found it is set on the context.
func SetCurrentUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid != nil {
			u := &models.User{}
			tx := c.Value("tx").(*pop.Connection)
			err := tx.Find(u, uid)
			if err != nil {
				if errors.Cause(err) == sql.ErrNoRows {
					return next(c)
				}
				return errors.WithStack(err)
			}
			c.Set("current_user", u)
		}
		return next(c)
	}
}
