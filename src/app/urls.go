package app

import (
	"github.com/rezwanul-haque/ID-Service/src/controllers/company"
	"github.com/rezwanul-haque/ID-Service/src/controllers/ping"
	"github.com/rezwanul-haque/ID-Service/src/controllers/users"
)

const (
	APIBase = "api/v1"
)

func mapUrls() {
	router.GET(APIBase+"/ping", ping.Ping)

	router.POST(APIBase+"/company/signup", company.CreateWithAdminUser)
	router.GET(APIBase+"/user/resolve", users.GetAll)
	router.POST(APIBase+"/user/signup", users.Create)
}
