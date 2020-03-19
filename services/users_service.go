package services

import (
	"github.com/rezwanul-haque/ID-Service/domain/users"
	"github.com/rezwanul-haque/ID-Service/utils/date"
	"github.com/rezwanul-haque/ID-Service/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct {
}

type usersServiceInterface interface {
	CreateAdminUser(users.User) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(apiKey string) (*users.User, *errors.RestErr)
	GetUserByCompanyIdAndRole(companyId int64, role string) (users.Users, *errors.RestErr)
}

func (u *usersService) CreateAdminUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	user.CreatedAt = date.GetNowDBFormat()
	user.UpdatedAt = date.GetNowDBFormat()

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *usersService) GetUser(apiKey string) (*users.User, *errors.RestErr) {
	result := &users.User{AppKey: apiKey}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (u *usersService) GetUserByCompanyIdAndRole(companyId int64, role string) (users.Users, *errors.RestErr) {
	result := &users.User{}
	return result.GetUsersByComapnyIdAndRole(companyId, role)
}
