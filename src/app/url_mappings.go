package app

import (
	"github.com/vadym-vasilyev/collector-go/src/controller/collector"
	"github.com/vadym-vasilyev/collector-go/src/controller/ping"

	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "github.com/vadym-vasilyev/collector-go/src/docs"
)

// @title Collector API
// @version 1.0
// @description Collector service API
// @termsOfService http://swagger.io/terms/

// @BasePath /api/v1
func mapUrls() {
	v1 := router.Group("/api/v1")
	v1.GET("/ping", ping.Ping)

	v1.POST("/collect/:app_token", collector.Collect)

	v1.GET("/internal/collect/:app_token", collector.Get)

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
