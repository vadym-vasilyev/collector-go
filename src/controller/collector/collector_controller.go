package collector

import (
	"github.com/gin-gonic/gin"
	"github.com/vadym-vasilyev/collector-go/src/domain/app_record"
	"github.com/vadym-vasilyev/collector-go/src/utils/logger"
	"github.com/vadym-vasilyev/collector-go/src/utils/rest_errors"
	"net/http"
)

// Collect godoc
// @Summary Add a batch of records
// @Description add by json account
// @Tags collect
// @Accept  json
// @Produce  json
// @Param app_token path string true "Application token"
// @Param batchRecords body app_record.RecordsBatch true "Add records batch"
// @Success 202
// @Failure 400 {object} rest_errors.RestErrStruct
// @Failure 404 {object} rest_errors.RestErrStruct
// @Failure 500 {object} rest_errors.RestErrStruct
// @Router /collect/{app_token} [post]
func Collect(c *gin.Context) {

	var recordsBatch app_record.RecordsBatch
	if err := c.ShouldBindJSON(&recordsBatch); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		logger.Error("Serialization error:", err)
		c.JSON(restErr.Status(), restErr)
		return
	}
	/*
		if _, saveErr := services.CollectorService.Save(&recordsBatch); saveErr != nil {
			c.JSON(saveErr.Status(), saveErr)
			return
		}
	*/
	c.Status(http.StatusAccepted)
}
