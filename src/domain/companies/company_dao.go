package companies

import (
	"github.com/rezwanul-haque/ID-Service/src/datasources/mysql/ids_db"
	"github.com/rezwanul-haque/ID-Service/src/logger"
	"github.com/rezwanul-haque/ID-Service/src/utils/errors"
)

const (
	queryInsertCompany = "INSERT INTO company(name, domain, created_at, updated_at) VALUES(?, ?, ?, ?);"
	queryGetCompany    = "SELECT name, domain FROM company WHERE id=?;"
)

func (company *Company) Save() *errors.RestErr {
	stmt, err := ids_db.Client.Prepare(queryInsertCompany)
	if err != nil {
		logger.Error("error when trying to prepare save company statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(company.Name, company.Domain, company.CreatedAt, company.UpdatedAt)

	if saveErr != nil {
		logger.Error("error when trying to save company", saveErr)
		return errors.NewInternalServerError("database error")
	}
	companyId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new company", err)
		return errors.NewInternalServerError("database error")
	}
	company.Id = companyId
	return nil
}

func (company *Company) Get() *errors.RestErr {
	stmt, err := ids_db.Client.Prepare(queryGetCompany)
	if err != nil {
		logger.Error("error when trying to prepare get company statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(company.Id)
	if getErr := result.Scan(&company.Name, &company.Domain); getErr != nil {
		logger.Error("error when trying to get company by id", getErr)
		return errors.NewInternalServerError("database error")
	}

	return nil
}
