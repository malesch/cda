package actions

import (
	"fmt"

	"github.com/cdacontrol/cda/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/pkg/errors"
)

// SceneEditorHandler ...
func SceneEditorHandler(c buffalo.Context) error {
	c.Set("sceneID", c.Param("scene"))

	return c.Render(200, r.HTML("scene-editor.html", "application.scene-editor.html"))
}

// SceneGetHandler ...
func SceneGetHandler(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	scene := &models.Scene{}
	if err := tx.Find(scene, c.Param("scene")); err != nil {
		c.Flash().Add("info", "Unknown Scene ID")
		return c.Redirect(302, "/")
	}

	sceneEvents, err := models.SceneEvents(tx, scene)
	if err != nil {
		c.Flash().Add("info", "No scene data available")
		return c.Error(404, err)
	}

	devices := models.Devices{}
	events := models.Events{}
	for d, es := range sceneEvents {
		devices = append(devices, d)
		for _, event := range es {
			events = append(events, event)
		}
	}

	return c.Render(200, r.JSON(map[string]interface{}{
		"devices": devices,
		"events":  events}))
}

// SceneUpdateHandler ...
func SceneUpdateHandler(c buffalo.Context) error {
	//tx := c.Value("tx").(*pop.Connection)
	c.Logger().Infof("Update => %+v", c.Param("scene"))

	return c.Render(200, r.JSON("OK"))
}

// SceneCreateHandler ...
func SceneCreateHandler(c buffalo.Context) error {
	scene := &models.Scene{}

	if err := c.Bind(scene); err != nil {
		return errors.WithStack(err)
	}

	tx := c.Value("tx").(*pop.Connection)

	verrs, err := tx.ValidateAndCreate(scene)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		c.Set("scene", scene)
		c.Set("errors", verrs)

		return c.Render(422, r.HTML("/scene/data"))
	}

	c.Flash().Add("success", "User was created successfully")

	return c.Render(200, r.JSON(scene))
}

// SceneDeleteHandler ...
func SceneDeleteHandler(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	scene := &models.Scene{}
	if err := tx.Find(scene, c.Param("scene")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(scene); err != nil {
		return errors.WithStack(err)
	}

	return c.Render(200, r.JSON(scene))
}

// -------
// Devices
// -------

// SceneDeviceCreateHandler ...
func SceneDeviceCreateHandler(c buffalo.Context) error {
	c.Logger().Infof("SceneDeviceCreateHandler: %v", c.Params())
	return c.Render(200, r.JSON("OK"))
}

// SceneDeviceUpdateHandler ...
func SceneDeviceUpdateHandler(c buffalo.Context) error {
	c.Logger().Infof("SceneDeviceUpdateHandler: %v", c.Params())
	return c.Render(200, r.JSON("OK"))
}

// SceneDeviceDeleteHandler ...
func SceneDeviceDeleteHandler(c buffalo.Context) error {
	c.Logger().Infof("SceneDeviceDeleteHandler: %v", c.Params())
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	sceneID := c.Param("scene")
	deviceID := c.Param("device")

	// Remove all device events
	events := models.Events{}
	query := tx.Where("scene_id = ? AND device_id = ?", sceneID, deviceID)
	err := query.All(&events)
	if err != nil {
		return c.Error(404, err)
	}
	for i := 0; i < len(events); i++ {
		event := events[i]
		tx.Destroy(&event)
	}

	// Redirect to the scenes index page
	return c.Render(200, r.JSON("OK"))
}

// ------
// Events
// ------

// SceneEventUpdateHandler ...
func SceneEventUpdateHandler(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	c.Logger().Infof("SceneEventUpdateHandler: %v", c.Params())

	event := models.Event{}
	if err := c.Bind(&event); err != nil {
		return errors.WithStack(err)
	}
	event.SceneID = uuid.FromStringOrNil(c.Param("scene"))
	c.Logger().Infof("SceneEventUpdateHandler Event: %v", event)

	if err := tx.Update(&event); err != nil {
		fmt.Printf("ERROR: %v", err)
		return c.Error(404, err)
	}
	return c.Render(200, r.JSON(event))
}

// SceneEventCreateHandler ...
func SceneEventCreateHandler(c buffalo.Context) error {
	c.Logger().Infof("SceneEventCreateHandler: %v", c.Params())
	tx := c.Value("tx").(*pop.Connection)
	event := models.Event{}
	if err := c.Bind(&event); err != nil {
		return errors.WithStack(err)
	}
	event.SceneID = uuid.FromStringOrNil(c.Param("scene"))
	c.Logger().Infof("SceneEventCreateHandler Event: %v", event)

	if err := tx.Create(&event); err != nil {
		return c.Error(404, err)
	}
	return c.Render(200, r.JSON(event))
}

// SceneEventDeleteHandler ...
func SceneEventDeleteHandler(c buffalo.Context) error {
	c.Logger().Infof("SceneEventDeleteHandler: %v", c.Params())
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Scene
	event := models.Event{}
	eventID := c.Param("event")
	c.Logger().Infof("eventID = %v", eventID)
	if err := tx.Find(&event, eventID); err != nil {
		return c.Error(404, err)
	}
	c.Logger().Infof("Retrievend Event: %v", event)

	c.Logger().Infof("1111 Event to delete: %v", event)
	if err := tx.Destroy(&event); err != nil {
		return errors.WithStack(err)
	}
	c.Logger().Infof("2222 Event returned : %v", event)
	// Redirect to the scenes index page
	return c.Render(200, r.JSON(event))
}
