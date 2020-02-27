package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// PingExample godoc
// @Summary ping example
// @Description do ping
// @Tags ping
// @Accept json
// @Produce json
// @Success 200 {string} string "pong"
// @Failure 500 {string} string "ok"
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
