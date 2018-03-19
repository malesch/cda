package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/pkg/errors"
)

type Scene struct {
	ID        uuid.UUID `json:"-" db:"id"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
	Name      string    `json:"name" db:"name"`
	MediumID  uuid.UUID `json:"-" db:"mediumID"`
}

// String is not required by pop and may be deleted
func (s Scene) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Scenes is not required by pop and may be deleted
type Scenes []Scene

// String is not required by pop and may be deleted
func (s Scenes) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *Scene) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: s.Name, Name: "Name"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *Scene) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *Scene) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// SceneEvents returns the scene events as a map events keyed by devices
func SceneEvents(tx *pop.Connection, scene *Scene) (map[Device][]Event, error) {
	events := []Event{}

	err := tx.Where("sceneID=?", scene.ID).All(&events)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// TODO:
	// HACK solution for stripping away attributes to be ignored for key
	// comparison e.g. CreatedAt or UpdatedAt. Better or idiomatic solution?
	res := make(map[Device][]Event)
	for _, event := range events {
		device := Device{}
		err := tx.Find(&device, event.DeviceID)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		props := Props{}
		err = tx.Where("eventID=?", event.ID).All(&props)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		event.Props = props

		deviceKey := Device{
			ID:   device.ID,
			Name: device.Name,
			Type: device.Type,
		}
		res[deviceKey] = append(res[deviceKey], event)
	}
	return res, nil
}
