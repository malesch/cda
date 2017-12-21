package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
)

type System struct {
	ID            uuid.UUID `json:"id" db:"id"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	IPAddress     string    `json:"ip_address" db:"ip_address"`
	XbeeGatewayID int       `json:"xbee_gateway_id" db:"xbee_gateway_id"`
	XbeeChannel   int       `json:"xbee_channel" db:"xbee_channel"`
	SystemTime    time.Time `json:"system_time" db:"-"`
}

// String is not required by pop and may be deleted
func (s System) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Systems is not required by pop and may be deleted
type Systems []System

// String is not required by pop and may be deleted
func (s Systems) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *System) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: s.IpAddress, Name: "IpAddress"},
		&validators.IntIsPresent{Field: s.XbeeGatewayID, Name: "XbeeGatewayID"},
		&validators.IntIsPresent{Field: s.XbeeChannel, Name: "XbeeChannel"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *System) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *System) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
