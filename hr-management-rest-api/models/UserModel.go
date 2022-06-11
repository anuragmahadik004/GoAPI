package models

type User struct {
	UserId   string
	UserName string
	Password string
	Email    string
	RoleInfo Role
}

type UserLogin struct {
	UserName string
	Password string
}

type Role struct {
	RoleId   string
	RoleName string
}
