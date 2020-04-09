package company

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rezwanul-haque/ID-Service/src/domain/companies"
	"github.com/rezwanul-haque/ID-Service/src/services"
	"github.com/rezwanul-haque/ID-Service/src/utils/consts"
	"github.com/rezwanul-haque/ID-Service/src/utils/errors"
	"github.com/rezwanul-haque/ID-Service/src/utils/helpers"
)

func Create(c *gin.Context) {
	secretKeyHeader := c.GetHeader("SecretKey")
	if helpers.IsInvalid(secretKeyHeader) || secretKeyHeader != consts.SecretKey {
		keyErr := errors.NewInternalServerError(fmt.Sprintf("secrectkey: '%s' is missing or invalid", secretKeyHeader))
		c.JSON(keyErr.Status, keyErr)
		return
	}
	var company companies.Company

	if err := c.ShouldBindJSON(&company); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CompanyService.CreateCompany(company)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func CreateWithAdminUser(c *gin.Context) {
	secretKeyHeader := c.GetHeader("SecretKey")
	if helpers.IsInvalid(secretKeyHeader) || secretKeyHeader != consts.SecretKey {
		keyErr := errors.NewInternalServerError(fmt.Sprintf("secrectkey: '%s' is missing or invalid", secretKeyHeader))
		c.JSON(keyErr.Status, keyErr)
		return
	}
	var company companies.CreateCompanyResponse

	if err := c.ShouldBindJSON(&company); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CompanyService.CreateCompanyWithAdminUser(company)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}
