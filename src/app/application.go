package app

import (
	"github.com/gin-gonic/gin"
	"github.com/vadym-vasilyev/collector-go/src/client/mongo_client"
	"github.com/vadym-vasilyev/collector-go/src/utils/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mongo_client.Init()

	mapUrls()
	logger.Info("about to start the application...")
	router.Run(":8080")
}
