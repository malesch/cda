package models

import "github.com/gobuffalo/buffalo"

type Role struct {
	ID   int
	Name string
}

var Roles = []Role{
	Role{ID: 0, Name: "user"},
	Role{ID: 1, Name: "admin"},
}

func SelectRoles(c buffalo.Context) map[string]int {
	roles := make(map[string]int)
	for _, role := range Roles {
		localizedName := Translate(c, "role."+role.Name)
		roles[localizedName] = role.ID
	}
	return roles
}

// func RoleName(id int) string {
// 	for _, role := range Roles {
// 		if role.ID == id {
// 			return role.Name
// 		}
// 	}
// 	return ""
// }
