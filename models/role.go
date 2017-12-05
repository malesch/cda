package models

var Roles = map[int]string{
	0: "user",
	1: "admin",
}

func SelectRoles() map[string]int {
	roles := make(map[string]int)
	for id, name := range Roles {
		roles["role."+name] = id
	}
	return roles
}
