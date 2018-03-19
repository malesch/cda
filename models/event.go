package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

// Event returns the data describing an event
type Event struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
	SceneID   uuid.UUID `json:"-" db:"scene_id"`
	DeviceID  uuid.UUID `json:"group" db:"device_id"`
	Start     int       `json:"start" db:"start_time"`
	End       int       `json:"end" db:"end_time"`
	Props     Props     `json:"props" db:"-"`
}

// String is not required by pop and may be deleted
func (e Event) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Events is not required by pop and may be deleted
type Events []Event

// String is not required by pop and may be deleted
func (e Events) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (e *Event) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: e.Start, Name: "Start"},
		&validators.IntIsPresent{Field: e.End, Name: "End"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (e *Event) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (e *Event) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// GetEventWithProps returns an Event by ID with all associated
func GetEventWithProps(tx *pop.Connection, id uuid.UUID) (*Event, error) {
	event := &Event{}
	if err := tx.Find(event, id); err != nil {
		return nil, fmt.Errorf("Cannot find Event ID %v", id)
	}

	props := Props{}
	if err := tx.Where("event_id=?", event.ID).All(&props); err != nil {
		return nil, fmt.Errorf("Error fetching props for Event ID %v", id)
	}
	event.Props = props
	return event, nil
}

// BeforeDestroy assures to remove also all associated Event properties
func (e *Event) BeforeDestroy(tx *pop.Connection) error {
	q := tx.RawQuery("delete from props where event_id=?", e.ID)
	return q.Exec()
}

// AfterCreate persists first the associated Event properties
func (e *Event) AfterCreate(tx *pop.Connection) error {
	if e.Props != nil {
		for _, prop := range e.Props {
			prop.EventID = e.ID
			if err := tx.Create(&prop); err != nil {
				return err
			}
		}
	}
	return nil
}

func mapifyProps(props Props) map[string]Prop {
	propMap := make(map[string]Prop)
	if props != nil {
		for _, prop := range props {
			propMap[prop.Name] = prop
		}
	}
	return propMap
}

// BeforeUpdate synchronizes the Props with the existing persisted Props
func (e *Event) BeforeUpdate(tx *pop.Connection) error {
	existingProps := Props{}
	if err := tx.Where("event_id = ?", e.ID).All(&existingProps); err != nil {
		return err
	}
	existingPropMap := mapifyProps(existingProps)
	newPropMap := mapifyProps(e.Props)

	for newPropName, newProp := range newPropMap {
		existingProp, ok := existingPropMap[newPropName]
		if ok {
			// update existing persisted value if altered
			if newProp.Value != existingProp.Value {
				existingProp.Value = newProp.Value
				if err := tx.Update(&existingProp); err != nil {
					return err
				}
			}
		} else {
			// add new property persisted value
			newProp.EventID = e.ID
			if err := tx.Create(&newProp); err != nil {
				return err
			}
		}
	}
	// Remove existing properties missing in the new properties
	for _, existingProp := range existingProps {
		if _, ok := newPropMap[existingProp.Name]; !ok {
			if err := tx.Destroy(&existingProp); err != nil {
				return err
			}
		}
	}

	return nil
}
