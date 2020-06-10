package users

import (
	"fmt"

	"github.com/rezwanul-haque/ID-Service/src/datasources/mysql/ids_db"
	"github.com/rezwanul-haque/ID-Service/src/logger"
	"github.com/rezwanul-haque/ID-Service/src/utils/errors"
)

const (
	queryInsertUser                 = "INSERT INTO user(company_id, user_id, app_key, role, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser                    = "SELECT user_id, app_key, company_id, role FROM user WHERE app_key=?;"
	queryGetUsersByCompanyIdAndRole = "SELECT id, user_id, app_key, company_id, role, created_at, updated_at FROM user WHERE company_id=? AND role=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := ids_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.AppKey)
	if getErr := result.Scan(&user.UserId, &user.AppKey, &user.CompanyId, &user.Role); getErr != nil {
		logger.Error("error when trying to get user by App Key", getErr)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := ids_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.CompanyId, user.UserId, user.AppKey, user.Role, user.CreatedAt, user.UpdatedAt)

	if saveErr != nil {
		logger.Error("error when trying to save user", saveErr)
		return errors.NewInternalServerError("database error")
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new user", err)
		return errors.NewInternalServerError("database error")
	}
	user.Id = userId
	return nil
}

func (user *User) GetUsersByComapnyIdAndRole(companyId int64, role string) ([]User, *errors.RestErr) {
	stmt, err := ids_db.Client.Prepare(queryGetUsersByCompanyIdAndRole)
	if err != nil {
		logger.Error("error when trying to prepare get users by companyId and role statement", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(companyId, role)
	if err != nil {
		logger.Error("error when trying to find users by status", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.UserId, &user.AppKey, &user.CompanyId, &user.Role, &user.CreatedAt, &user.UpdatedAt); err != nil {
			logger.Error("error when scan user row into user struct", err)
			return nil, errors.NewInternalServerError("database error")
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching company id %d and role %s", companyId, role))
	}
	return results, nil
}
