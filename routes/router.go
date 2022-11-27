package routes

import (
	"backup/routes/endpoints"
	"github.com/gin-gonic/gin"
)

type Router interface {
	Route(engine *gin.Engine)
}

type Loader struct {
}

func (loader Loader) Load() []Router {
	healthcheck := new(endpoints.HealthCheck)
	backup := new(endpoints.Backup)
	return []Router{
		healthcheck,
		backup,
	}
}
