package actions

import (
	"github.com/cdacontrol/cda/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
)

// SceneEditorHandler ...
func SceneEditorHandler(c buffalo.Context) error {

	c.Set("scene_id", c.Param("scene_id"))

	return c.Render(200, r.HTML("scene-editor.html", "application.scene-editor.html"))
}

// SceneDataHandler ...
func SceneDataHandler(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	scene := &models.Scene{}
	if err := tx.Find(scene, c.Param("scene_id")); err != nil {
		c.Flash().Add("info", "Unknown Scene ID")
		return c.Redirect(302, "/")
	}

	sceneEvents, err := models.SceneEvents(tx, scene)
	if err != nil {
		c.Flash().Add("info", "No scene data available")
		return c.Error(404, err)
	}

	c.Logger().Infof("Scene ID = %s", scene.ID)
	c.Logger().Infof("SceneEvents = %+v", sceneEvents)

	type DeviceData struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"content"`
		Type string    `json:"type"`
	}
	type EventData struct {
		ID       uuid.UUID         `json:"id"`
		DeviceID uuid.UUID         `json:"group"`
		Start    int               `json:"start"`
		End      int               `json:"end"`
		Props    map[string]string `json:"props"`
	}
	devices := []DeviceData{}
	events := []EventData{}

	for d, e := range sceneEvents {
		devices = append(devices, DeviceData{
			ID:   d.ID,
			Name: d.Name,
			Type: d.Type,
		})
		for _, event := range e {
			props := make(map[string]string)
			for _, prop := range event.Props {
				props[prop.Name] = prop.Value
			}
			events = append(events, EventData{
				ID:       event.ID,
				DeviceID: d.ID,
				Start:    event.Start,
				End:      event.End,
				Props:    props,
			})
		}
	}

	return c.Render(200, r.JSON(map[string]interface{}{
		"devices": devices,
		"events":  events}))
}
