package routes

import (
	"backup/cfg"
	"backup/routes/endpoints"
	"github.com/gin-gonic/gin"
)

type Router interface {
	Route(engine *gin.Engine)
}

type Loader struct {
}

func (loader Loader) Load(conf *cfg.Config) []Router {
	healthcheck := new(endpoints.HealthCheck)
	backup := &endpoints.Backup{
		Conf: conf,
	}
	return []Router{
		healthcheck,
		backup,
	}
}
