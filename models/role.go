package models

// Roles is the map of available user roles
var Roles = map[int]string{
	1: "user",
	2: "admin",
}

// SelectRoles returns a map of Role localization ID's with
// their associated role number
func SelectRoles() map[string]int {
	roles := make(map[string]int)
	for id, name := range Roles {
		roles["role."+name] = id
	}
	return roles
}
