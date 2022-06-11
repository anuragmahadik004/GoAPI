package interfaces

import (
	"github.com/anuragmahadik004/hr_api/datalayer"
	"github.com/anuragmahadik004/hr_api/models"
)

type UserRepository struct {
}

var dlUser datalayer.DLUser

func (UserRepository) SaveUser(User models.User) bool {
	return dlUser.SaveUser(User)
}

func (UserRepository) GetUser(UserName string) models.User {
	return dlUser.GetUser(UserName)
}
