package models_test

import (
	"fmt"
	"reflect"

	"github.com/cdacontrol/cda/models"
)

func (ms *ModelSuite) Test_User_Create() {
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	u := &models.User{
		Name:                 "John Doe",
		Password:             "password",
		PasswordConfirmation: "password",
		LocaleID:             1,
		RoleID:               1,
	}
	ms.Zero(u.PasswordHash, "PasswordHash should not be manually set")

	verrs, err := ms.DB.ValidateAndCreate(u)
	ms.NoError(err)
	ms.False(verrs.HasAny(), fmt.Sprintf("Error messages:\n %+v", verrs))
	ms.NotZero(u.PasswordHash, "PasswordHash is missing")

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)
}

func (ms *ModelSuite) Test_User_Create_ValidationErrors() {
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	u := &models.User{
		Name:     "John Doe",
		Password: "password",
	}
	ms.Zero(u.PasswordHash, "PasswordHash should not be manually set")

	verrs, err := ms.DB.ValidateAndCreate(u)
	ms.NoError(err)
	ms.True(verrs.HasAny())
	ms.Equal(3, len(verrs.Errors),
		fmt.Sprintf("There should be exactly 3 validation errors [role_id locale_id password_confirmation]\nbut only got %v",
			reflect.ValueOf(verrs.Errors).MapKeys()))

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)
}

func (ms *ModelSuite) Test_User_Create_UserExists() {
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	u := &models.User{
		Name:                 "John Doe",
		Password:             "password",
		PasswordConfirmation: "password",
		LocaleID:             1,
		RoleID:               1,
	}
	ms.Zero(u.PasswordHash, "PasswordHash should not be manually set")
	verrs, err := ms.DB.ValidateAndCreate(u)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	ms.NotZero(u.PasswordHash, "PasswordHash is missing")

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)

	u = &models.User{
		Name:                 "John Doe",
		Password:             "password",
		PasswordConfirmation: "password",
		LocaleID:             1,
		RoleID:               1,
	}
	verrs, err = ms.DB.ValidateAndCreate(u)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)
}

func (ms *ModelSuite) Test_User_Update() {
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	u := &models.User{
		Name:                 "John Doe",
		Password:             "password",
		PasswordConfirmation: "password",
		LocaleID:             1,
		RoleID:               1,
	}
	ms.Zero(u.PasswordHash, "PasswordHash should not be manually set")
	verrs, err := ms.DB.ValidateAndCreate(u)

	passwordHash_before := u.PasswordHash

	ms.NoError(err)
	ms.False(verrs.HasAny())
	ms.NotZero(u.PasswordHash, "PasswordHash is missing")

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)

	u.Password = "password2"
	u.PasswordConfirmation = "password2"
	u.LocaleID = 3
	u.RoleID = 2

	verrs, err = ms.DB.ValidateAndUpdate(u)
	ms.NoError(err)
	ms.False(verrs.HasAny(), "Updating user failed")

	// Read back user
	u_updated := &models.User{}

	err = ms.DB.Find(u_updated, u.ID)
	ms.NoError(err, "Reading user failed")

	ms.Equal(u.ID, u_updated.ID)
	ms.NotEqual(passwordHash_before, u_updated.PasswordHash)
	ms.NotEqual(1, u_updated.LocaleID)
	ms.NotEqual(1, u_updated.RoleID)
}

func (ms *ModelSuite) Test_User_Destroy() {
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	u := &models.User{
		Name:                 "John Doe",
		Password:             "password",
		PasswordConfirmation: "password",
		LocaleID:             1,
		RoleID:               1,
	}
	ms.Zero(u.PasswordHash, "PasswordHash should not be manually set")
	verrs, err := ms.DB.ValidateAndCreate(u)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	ms.NotZero(u.PasswordHash, "PasswordHash is missing")

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)

	err = ms.DB.Destroy(u)
	ms.NoError(err)

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)
}
