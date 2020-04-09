package ids_db

import (
	"database/sql"
	"fmt"
	"github.com/rezwanul-haque/ID-Service/src/logger"
	"github.com/rezwanul-haque/ID-Service/src/utils/helpers"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB

	username = helpers.GoDotEnvVariable("MYSQL_IDS_USERNAME")
	password = helpers.GoDotEnvVariable("MYSQL_IDS_PASSWORD")
	host     = helpers.GoDotEnvVariable("MYSQL_IDS_HOST")
	port     = helpers.GoDotEnvVariable("MYSQL_IDS_PORT")
	schema   = helpers.GoDotEnvVariable("MYSQL_IDS_SCHEMA")
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		username, password, host, port, schema,
	)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		logger.Error("connecting to database failed: ", err)
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	logger.Info("database successfully configured")
}
