package app

import (
	"github.com/vadym-vasilyev/collector-go/src/controller/collector"
	"github.com/vadym-vasilyev/collector-go/src/controller/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/collect/:app_token", collector.Collect)
}
