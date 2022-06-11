package models

type LoginSession struct {
	UserInfo       User
	SessionTimeOut string
	SessionId      string
	SessionEnd     bool
}
