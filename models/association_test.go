package models_test

import (
	"fmt"

	"github.com/gobuffalo/pop"
	"github.com/malesch/cda/models"
	"github.com/pkg/errors"
)

// Helper fn

func createScene(tx *pop.Connection, name string) (*models.Scene, error) {
	scene := &models.Scene{Name: name}
	if err := tx.Create(scene); err != nil {
		return nil, errors.WithStack(err)
	}
	return scene, nil
}

func createEvent(tx *pop.Connection, scene *models.Scene, device *models.Device, start int, end int) (*models.Event, error) {
	event := &models.Event{SceneID: scene.ID, DeviceID: device.ID, Start: start, End: end}
	if err := tx.Create(event); err != nil {
		return nil, err
	}
	return event, nil
}

func createDevice(tx *pop.Connection, name string, typ string) (*models.Device, error) {
	device := &models.Device{Name: name, Type: typ}
	if err := tx.Create(device); err != nil {
		return nil, err
	}
	return device, nil
}

func createProp(tx *pop.Connection, event *models.Event, name string, value string) (*models.Prop, error) {
	prop := &models.Prop{EventID: event.ID, Name: name, Value: value}
	if err := tx.Create(prop); err != nil {
		return nil, err
	}
	return prop, nil
}

func populateDB(tx *pop.Connection) error {
	scene, err := createScene(tx, "SceneTest")
	if err != nil {
		return errors.WithStack(err)
	}
	device, err := createDevice(tx, "DeviceTest", "Light")
	if err != nil {
		return errors.WithStack(err)
	}
	event, err := createEvent(tx, scene, device, 12, 42)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = createProp(tx, event, "TestPropName0", "TestPropValue0")
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = createProp(tx, event, "TestPropName1", "TestPropValue1")
	return err
}

func (ms *ModelSuite) Test_Association_Loading() {
	err := populateDB(ms.DB)
	ms.NoError(err)

	// load eagerly Scene
	scene := &models.Scene{}
	err = ms.DB.Eager().First(scene)
	ms.NoError(err)
	//fmt.Printf("### Scene -> %v\n", scene)

	events := scene.Events
	ms.NotEmpty(events)
	ms.Len(events, 1)

	event := &events[0]
	ms.DB.Load(event)
	//fmt.Printf("### Event -> %v\n", event)

	ms.Equal(event.Device.Name, "DeviceTest")
	ms.Equal(event.Device.Type, "Light")

	props := event.Props
	ms.NotEmpty(props)
	ms.Len(props, 2)

	for i := 0; i < 2; i++ {
		prop := props[i]

		//fmt.Printf("### Prop %d -> %v\n", i, prop)
		ms.Equal(prop.Name, fmt.Sprintf("TestPropName%d", i))
		ms.Equal(prop.Value, fmt.Sprintf("TestPropValue%d", i))
	}
}
