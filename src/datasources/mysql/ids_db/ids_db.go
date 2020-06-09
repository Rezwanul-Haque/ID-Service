package ids_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

	"github.com/rezwanul-haque/ID-Service/src/logger"
	"os"
)

var (
	Client *sql.DB

	username = os.Getenv("MYSQL_IDS_USERNAME")
	password = os.Getenv("MYSQL_IDS_PASSWORD")
	host     = os.Getenv("MYSQL_IDS_HOST")
	port     = os.Getenv("MYSQL_IDS_PORT")
	schema   = os.Getenv("MYSQL_IDS_SCHEMA")
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
