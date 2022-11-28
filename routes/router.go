package routes

import (
	"backup/cfg"
	"backup/routes/endpoints"
	"backup/rsync"
	"github.com/gin-gonic/gin"
)

type Router interface {
	Route(engine *gin.Engine)
}

type Loader struct {
}

func (loader Loader) Load(conf *cfg.Config, rsyncExecutor *rsync.Executor) []Router {
	healthcheck := new(endpoints.HealthCheck)
	backup := &endpoints.Backup{
		Conf:          conf,
		RSyncExecutor: rsyncExecutor,
	}
	return []Router{
		healthcheck,
		backup,
	}
}
