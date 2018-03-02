package actions

import (
	"github.com/cdacontrol/cda/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
)

// SystemsResource is the resource for the System model
type SystemsResource struct {
	buffalo.Resource
}

// List renders the Edit page
func (v SystemsResource) List(c buffalo.Context) error {
	return v.Edit(c)
}

// Show renders the Edit page
func (v SystemsResource) Show(c buffalo.Context) error {
	return v.Edit(c)
}

// New renders the Edit page
func (v SystemsResource) New(c buffalo.Context) error {
	return v.Edit(c)
}

// Create adds a System to the DB. This function is mapped to the
// path POST /systems
func (v SystemsResource) Create(c buffalo.Context) error {
	system := &models.System{}

	if err := c.Bind(system); err != nil {
		return errors.WithStack(err)
	}

	tx := c.Value("tx").(*pop.Connection)

	verrs, err := tx.ValidateAndCreate(system)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		c.Set("system", system)
		c.Set("errors", verrs)
		return c.Render(422, r.HTML("systems/new.html"))
	}

	c.Flash().Add("success", "System was created successfully")

	return c.Redirect(302, "/systems/%s", system.ID)
}

// Edit renders a edit form for a System. This function is
// mapped to the path GET /systems/{system_id}/edit
func (v SystemsResource) Edit(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	system := &models.System{}

	if err := tx.First(system); err != nil {
		// Read System defaults e.g. from file sysstem
		system.IPAddress = "0.0.0.0"
		system.XbeeChannel = -1
		system.XbeeGatewayID = -1

		verrs, err := tx.ValidateAndCreate(system)
		if err != nil {
			return errors.WithStack(err)
		}

		if verrs.HasAny() {
			c.Set("errors", verrs)
		}
	}

	c.Set("system", system)
	return c.Render(200, r.HTML("systems/edit.html"))
}

// Update changes a System in the DB. This function is mapped to
// the path PUT /systems/0 and can *only* update  the first entry
func (v SystemsResource) Update(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	system := &models.System{}

	if err := tx.First(system); err != nil {
		return c.Error(404, err)
	}

	if err := c.Bind(system); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(system)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		c.Set("system", system)
		c.Set("errors", verrs)

		return c.Render(422, r.HTML("systems/edit.html"))
	}

	c.Flash().Add("success", "System was updated successfully")

	return c.Redirect(302, "/")
}
