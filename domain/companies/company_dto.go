package companies

import (
	"strings"

	"github.com/rezwanul-haque/ID-Service/domain/users"
	"github.com/rezwanul-haque/ID-Service/utils/errors"
)

type Company struct {
	Id        int64  `json:"id"`
	Domain    string `json:"domain"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type CreateCompanyResponse struct {
	Domain string     `json:"domain"`
	Name   string     `json:"name"`
	Admin  users.User `json:"admin"`
}

func (company *CreateCompanyResponse) TrimRequestBody() *errors.RestErr {
	company.Domain = strings.TrimSpace(company.Domain)
	company.Name = strings.TrimSpace(company.Name)

	return nil
}
