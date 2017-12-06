package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
)

type User struct {
	ID                   uuid.UUID `json:"id" db:"id"`
	CreatedAt            time.Time `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time `json:"updated_at" db:"updated_at"`
	Name                 string    `json:"name" db:"name"`
	Password             string    `json:"-" db:"-"`
	PasswordConfirmation string    `json:"-" db:"-"`
	PasswordHash         string    `json:"-" db:"password_hash"`
	LocaleID             int       `json:"locale_id" db:"locale_id"`
	RoleID               int       `json:"role_id" db:"role_id"`
}

// String is not required by pop and may be deleted
func (u User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Users is not required by pop and may be deleted
type Users []User

// String is not required by pop and may be deleted
func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Name, Name: "Name"},
		&validators.StringIsPresent{Field: u.Password, Name: "Password"},
		&validators.StringIsPresent{Field: u.PasswordConfirmation, Name: "PasswordConfirmation"},
		//&validators.StringIsPresent{Field: u.PasswordHash, Name: "PasswordHash"},
		&validators.IntIsPresent{Field: u.LocaleID, Name: "Locale"},
		&validators.IntIsPresent{Field: u.RoleID, Name: "Role"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// BeforeSave is called before storing the entity to the database and makes sure that
// the password hash is updated based on the (confirmed) password value
func (u *User) BeforeSave(tx *pop.Connection) error {
	if u.Password != "" {
		ph, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return errors.WithStack(err)
		}
		u.PasswordHash = string(ph)
	}
	return nil
}
