package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthCheck struct {
}

func (hc HealthCheck) Route(engine *gin.Engine) {
	engine.GET("/healthcheck", func(context *gin.Context) {
		context.String(http.StatusOK, "")
	})
}
