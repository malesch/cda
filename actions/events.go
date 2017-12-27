package actions

import (
	"github.com/cdacontrol/cda/models"
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Event)
// DB Table: Plural (events)
// Resource: Plural (Events)
// Path: Plural (/events)
// View Template Folder: Plural (/templates/events/)

// EventsResource is the resource for the Event model
type EventsResource struct {
	buffalo.Resource
}

// List gets all Events. This function is mapped to the path
// GET /events
func (v EventsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	events := &models.Events{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Events from the DB
	if err := q.All(events); err != nil {
		return errors.WithStack(err)
	}

	// Make Events available inside the html template
	c.Set("events", events)

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.HTML("events/index.html"))
}

// Show gets the data for one Event. This function is mapped to
// the path GET /events/{event_id}
func (v EventsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Event
	event := &models.Event{}

	// To find the Event the parameter event_id is used.
	if err := tx.Find(event, c.Param("event_id")); err != nil {
		return c.Error(404, err)
	}

	// Make event available inside the html template
	c.Set("event", event)

	return c.Render(200, r.HTML("events/show.html"))
}

// New renders the form for creating a new Event.
// This function is mapped to the path GET /events/new
func (v EventsResource) New(c buffalo.Context) error {
	// Make event available inside the html template
	c.Set("event", &models.Event{})

	return c.Render(200, r.HTML("events/new.html"))
}

// Create adds a Event to the DB. This function is mapped to the
// path POST /events
func (v EventsResource) Create(c buffalo.Context) error {
	// Allocate an empty Event
	event := &models.Event{}

	// Bind event to the html form elements
	if err := c.Bind(event); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(event)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make event available inside the html template
		c.Set("event", event)

		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("events/new.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Event was created successfully")

	// and redirect to the events index page
	return c.Redirect(302, "/events/%s", event.ID)
}

// Edit renders a edit form for a Event. This function is
// mapped to the path GET /events/{event_id}/edit
func (v EventsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Event
	event := &models.Event{}

	if err := tx.Find(event, c.Param("event_id")); err != nil {
		return c.Error(404, err)
	}

	// Make event available inside the html template
	c.Set("event", event)
	return c.Render(200, r.HTML("events/edit.html"))
}

// Update changes a Event in the DB. This function is mapped to
// the path PUT /events/{event_id}
func (v EventsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Event
	event := &models.Event{}

	if err := tx.Find(event, c.Param("event_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Event to the html form elements
	if err := c.Bind(event); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(event)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make event available inside the html template
		c.Set("event", event)

		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("events/edit.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Event was updated successfully")

	// and redirect to the events index page
	return c.Redirect(302, "/events/%s", event.ID)
}

// Destroy deletes a Event from the DB. This function is mapped
// to the path DELETE /events/{event_id}
func (v EventsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Event
	event := &models.Event{}

	// To find the Event the parameter event_id is used.
	if err := tx.Find(event, c.Param("event_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(event); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "Event was destroyed successfully")

	// Redirect to the events index page
	return c.Redirect(302, "/events")
}