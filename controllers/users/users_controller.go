package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rezwanul-haque/ID-Service/domain/users"
	"github.com/rezwanul-haque/ID-Service/services"
	"github.com/rezwanul-haque/ID-Service/utils/consts"
	"github.com/rezwanul-haque/ID-Service/utils/errors"
	"github.com/rezwanul-haque/ID-Service/utils/helpers"
)

func Create(c *gin.Context) {
	appKeyHeader := c.GetHeader("AppKey")
	if helpers.IsInvalid(appKeyHeader) {
		keyErr := errors.NewInternalServerError(fmt.Sprintf("Appkey: '%s' is missing", appKeyHeader))
		c.JSON(keyErr.Status, keyErr)
		return
	}

	foundUser, getErr := services.UsersService.GetUser(appKeyHeader)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	if appKeyHeader != foundUser.AppKey {
		keyErr := errors.NewInternalServerError(fmt.Sprintf("Appkey: '%s' is invalid", appKeyHeader))
		c.JSON(keyErr.Status, keyErr)
		return
	}

	var user users.User
	user.CompanyId = foundUser.CompanyId
	user.Role = consts.Role.USER

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetAll(c *gin.Context) {
	appKeyHeader := c.GetHeader("AppKey")
	if helpers.IsInvalid(appKeyHeader) {
		keyErr := errors.NewInternalServerError(fmt.Sprintf("Appkey: '%s' is missing", appKeyHeader))
		c.JSON(keyErr.Status, keyErr)
		return
	}

	foundUser, getErr := services.UsersService.GetUser(appKeyHeader)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	if appKeyHeader != foundUser.AppKey {
		keyErr := errors.NewInternalServerError(fmt.Sprintf("Appkey: '%s' is invalid", appKeyHeader))
		c.JSON(keyErr.Status, keyErr)
		return
	}

	var result users.ResolveResponse

	company, getErr := services.CompanyService.GetCompany(foundUser.CompanyId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	result.CompanuName = company.Name
	result.Domain = company.Domain
	result.CompanyId = foundUser.CompanyId
	result.UserId = foundUser.UserId
	result.Role = foundUser.Role

	subordinates, getErr := services.UsersService.GetUserByCompanyIdAndRole(foundUser.CompanyId, consts.Role.USER)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	result.Subordinates = subordinates

	c.JSON(http.StatusOK, result)
}
