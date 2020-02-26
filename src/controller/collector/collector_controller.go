package collector

import (
	"github.com/gin-gonic/gin"
	"github.com/vadym-vasilyev/collector-go/src/domain/app_record"
	"github.com/vadym-vasilyev/collector-go/src/services"
	"github.com/vadym-vasilyev/collector-go/src/utils/rest_errors"
	"net/http"
)

func Collect(c *gin.Context) {

	var recordsBatch app_record.RecordsBatch
	if err := c.ShouldBindJSON(&recordsBatch); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	if _, saveErr := services.CollectorService.Save(&recordsBatch); saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}

	c.Status(http.StatusAccepted)
}
