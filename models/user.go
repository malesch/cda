package models

import (
	"encoding/json"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

// User model
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

// PasswordConfirmationValidator holds the data for performing the validation
// of the password confirmation
type PasswordConfirmationValidator struct {
	Password             string
	PasswordConfirmation string
}

// IsValid implements the check for identical Password and PasswordConfirmation
// if a password is available
func (u *PasswordConfirmationValidator) IsValid(errors *validate.Errors) {
	if u.Password != "" {
		if u.Password != u.PasswordConfirmation {
			errors.Add(validators.GenerateKey("password_confirmation"), "Password and PasswordConfirmation are not identical")
		}
	}
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Name, Name: "Name"},
		// check to see if the user name is already taken:
		&validators.FuncValidator{
			Field:   u.Name,
			Name:    "Name",
			Message: "%s is already taken",
			Fn: func() bool {
				var b bool
				q := tx.Where("name = ?", u.Name)
				if u.ID != uuid.Nil {
					q = q.Where("id != ?", u.ID)
				}
				b, err = q.Exists(u)
				if err != nil {
					return false
				}
				return !b
			},
		},
		&PasswordConfirmationValidator{Password: u.Password, PasswordConfirmation: u.PasswordConfirmation},
		&validators.IntIsGreaterThan{Name: "LocaleID", Field: u.LocaleID, Compared: 0},
		&validators.IntIsGreaterThan{Name: "RoleID", Field: u.RoleID, Compared: 0},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Password, Name: "Password"},
		&validators.StringIsPresent{Field: u.PasswordConfirmation, Name: "PasswordConfirmation"},
	), err
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
