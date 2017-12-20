package models

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/gobuffalo/envy"
	"github.com/markbates/pop"
	"github.com/markbates/pop/nulls"
	"github.com/markbates/validate"
	"github.com/satori/go.uuid"
)

type Medium struct {
	ID        uuid.UUID    `json:"id" db:"id"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	Name      nulls.String `json:"name" db:"name"`
	Type      string       `json:"type" db:"type"`
	Size      int          `json:"size" db:"size"`
	FileID    string       `json:"fileId" db:"-"`
	FileName  string       `json:"fileName" db:"-"`
}

// String is not required by pop and may be deleted
func (m Medium) String() string {
	jm, _ := json.Marshal(m)
	return string(jm)
}

// Media is not required by pop and may be deleted
type Media []Medium

// String is not required by pop and may be deleted
func (m Media) String() string {
	jm, _ := json.Marshal(m)
	return string(jm)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (m *Medium) Validate(tx *pop.Connection) (*validate.Errors, error) {
	fmt.Printf("Validate!\n")
	return validate.Validate(
	//		&validators.StringIsPresent{Field: m.Type, Name: "Type"},
	//		&validators.IntIsPresent{Field: m.Size, Name: "Size"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (m *Medium) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (m *Medium) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// AfterCreate is called after persisting the Medium record and copies the media file
// from the temporary upload location to its final destination and renames it to the
// UUID of the Medium record.
func (m *Medium) AfterCreate(tx *pop.Connection) error {
	fmt.Printf("Move and rename Medium file.\n")
	infile, err := os.Open(filepath.Join(os.TempDir(), "_cda", m.FileID))
	if err != nil {
		return err
	}

	destPath, err := MediaStorePath()
	fmt.Printf("Media destPath = %s\n", destPath)
	if err != nil {
		return err
	}

	outfile, err := os.Create(filepath.Join(destPath, m.ID.String()))
	if err != nil {
		return err
	}
	_, err = io.Copy(outfile, infile)
	return err
}

func MediaStorePath() (string, error) {
	path := envy.Get("mediaPath", "./media/")
	stat, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return "", err
			}
			return path, nil
		}
		return "", err
	}
	if stat.IsDir() {
		return path, nil
	}
	return "", fmt.Errorf("Cannot create directory `%s` because it is a regular file", path)
}
