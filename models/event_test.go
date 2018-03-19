package models_test

import (
	"fmt"

	"github.com/cdacontrol/cda/models"
)

func (ms *ModelSuite) Test_Event_Create() {
	count, err := ms.DB.Count("scenes")
	ms.NoError(err)
	ms.Equal(0, count)

	scene := &models.Scene{
		Name: "Test Scene",
	}
	err = ms.DB.Create(scene)
	ms.NoError(err)

	light1 := &models.Device{
		Name: "Osram Leuchtband",
		Type: "Light",
	}
	err = ms.DB.Create(light1)
	ms.NoError(err)

	e1 := &models.Event{
		SceneID:  scene.ID,
		DeviceID: light1.ID,
		Start:    14400000,
		End:      18960000,
	}
	err = ms.DB.Create(e1)
	ms.NoError(err)

	err = ms.DB.Create(&models.Prop{
		EventID: e1.ID,
		Name:    "color",
		Value:   "red",
	})
	ms.NoError(err)

	err = ms.DB.Create(&models.Prop{
		EventID: e1.ID,
		Name:    "saturation",
		Value:   "159",
	})

	s := models.Scenes{}
	ms.NoError(err)

	for _, sc := range s {
		fmt.Printf("Scene %#v\n", sc)
	}

	// Load event with props (manually...)
	event, err := models.GetEventWithProps(ms.DB, e1.ID)
	ms.NoError(err)

	props := &models.Props{}
	err = ms.DB.Where("event_id = ?", event.ID).All(props)
	ms.NoError(err)

	// change a prop
	(*props)[1].Value = "3.141592"

	// add a prop
	newProps := append(*props, models.Prop{Name: "a_new_prop", Value: "COOOL!"})

	// remove a prop
	event.Props = append(newProps[:0], newProps[1:]...)

	ms.DB.Update(event)
}
