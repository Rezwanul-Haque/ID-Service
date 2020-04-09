package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rezwanul-haque/ID-Service/src/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("about to start the application")
	router.Run(":8080")
}
