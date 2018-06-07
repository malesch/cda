package main

import (
	"fmt"
	"log"

	"github.com/gobuffalo/pop"
	"github.com/malesch/cda/models"
	"github.com/pkg/errors"
)

// DB connection
var DB *pop.Connection

func init() {
	var err error
	DB, err = pop.Connect("development")
	if err != nil {
		log.Fatal(err)
	}
	//pop.Debug = true
}

func main() {
	if err := DB.TruncateAll(); err != nil {
		log.Fatal(err)
	}
	if err := populateDB(); err != nil {
		log.Fatal(err)
	}
	loadScene(DB)
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

func populateDB() error {
	return DB.Transaction(func(tx *pop.Connection) error {
		device, err := createDevice(tx, "Device 1", "Device Type 1")
		if err != nil {
			return errors.WithStack(err)
		}

		scene := &models.Scene{Name: "Test scene"}
		if err := tx.Create(scene); err != nil {
			return errors.WithStack(err)
		}

		fmt.Println(scene.ID, device.ID)
		event1 := &models.Event{SceneID: scene.ID, DeviceID: device.ID, Start: 1, End: 2}
		if err := tx.Create(event1); err != nil {
			return errors.WithStack(err)
		}

		fmt.Println(scene.ID, device.ID)
		event2 := &models.Event{SceneID: scene.ID, DeviceID: device.ID, Start: 3, End: 4}
		if err := tx.Create(event2); err != nil {
			return errors.WithStack(err)
		}

		_, err = createProp(tx, event1, "Prop1", "Value1")
		if err != nil {
			return errors.WithStack(err)
		}

		_, err = createProp(tx, event1, "Prop2", "Value2")
		if err != nil {
			return errors.WithStack(err)
		}

		_, err = createProp(tx, event1, "Prop3", "Value3")
		if err != nil {
			return errors.WithStack(err)
		}

		return nil
	})
}

func loadScene(tx *pop.Connection) error {
	fmt.Printf("\n\n### --- LOAD SCENE --- ###\n\n")
	s := &models.Scene{}
	if err := DB.Eager().First(s); err != nil {
		return errors.WithStack(err)
	}
	fmt.Println("### s -> ", s)
	events := s.Events
	fmt.Printf("There are %d events\n", len(s.Events))
	for _, e := range events {
		err := tx.Load(&e)
		if err != nil {
			return errors.WithStack(err)
		}
		fmt.Printf("Event: %v\n", e)
		fmt.Println("### e.Props -> ", e.Props)
	}

	return nil
}

