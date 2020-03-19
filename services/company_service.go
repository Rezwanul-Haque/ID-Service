package services

import (
	guuid "github.com/google/uuid"
	"github.com/rezwanul-haque/ID-Service/domain/companies"
	"github.com/rezwanul-haque/ID-Service/domain/users"
	"github.com/rezwanul-haque/ID-Service/utils/consts"
	"github.com/rezwanul-haque/ID-Service/utils/date"
	"github.com/rezwanul-haque/ID-Service/utils/errors"
	"github.com/rezwanul-haque/ID-Service/utils/hash"
)

var (
	CompanyService companyServiceInterface = &companyService{}
)

type companyService struct {
}

type companyServiceInterface interface {
	CreateCompany(companies.Company) (*companies.Company, *errors.RestErr)
	CreateCompanyWithAdminUser(companies.CreateCompanyResponse) (*companies.CreateCompanyResponse, *errors.RestErr)
	GetCompany(companyId int64) (*companies.Company, *errors.RestErr)
}

func (c *companyService) CreateCompany(company companies.Company) (*companies.Company, *errors.RestErr) {
	if err := company.Save(); err != nil {
		return nil, err
	}
	return &company, nil
}

func (c *companyService) CreateCompanyWithAdminUser(company companies.CreateCompanyResponse) (*companies.CreateCompanyResponse, *errors.RestErr) {
	if err := company.TrimRequestBody(); err != nil {
		return nil, err
	}

	createdAt := date.GetNowDBFormat()
	updatedAt := date.GetNowDBFormat()

	var companyObj companies.Company

	companyObj.Domain = company.Domain
	companyObj.Name = company.Name
	companyObj.CreatedAt = createdAt
	companyObj.UpdatedAt = updatedAt

	companyResult, createErr := CompanyService.CreateCompany(companyObj)
	if createErr != nil {
		return nil, createErr
	}

	var user users.User
	user.CompanyId = companyResult.Id
	user.UserId = guuid.New().String()
	user.AppKey = hash.GetMD5Hash(company.Domain)
	user.Role = consts.Role.ADMIN
	user.CreatedAt = createdAt
	user.UpdatedAt = updatedAt

	userResult, createErr := UsersService.CreateAdminUser(user)
	if createErr != nil {
		return nil, createErr
	}

	company.Admin = *userResult
	return &company, nil
}

func (c *companyService) GetCompany(companyId int64) (*companies.Company, *errors.RestErr) {
	result := &companies.Company{Id: companyId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}
