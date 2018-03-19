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
	MediumID  uuid.UUID `json:"-" db:"medium_id"`
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
// TODO: Refactor query (-> left join query), inefficient implementation!
func SceneEvents(tx *pop.Connection, scene *Scene) (map[Device][]Event, error) {
	events := []Event{}

	err := tx.Where("scene_id=?", scene.ID).All(&events)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	res := make(map[Device][]Event)

	devices := []Device{}
	err = tx.All(&devices)

	if err != nil {
		return nil, errors.WithStack(err)
	}
	for _, device := range devices {
		events := []Event{}
		err = tx.Where("device_id=?", device.ID).All(&events)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		for idx, event := range events {
			props := Props{}
			err = tx.Where("event_id=?", event.ID).All(&props)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			events[idx].Props = props
		}

		res[device] = events
	}

	return res, nil
}
