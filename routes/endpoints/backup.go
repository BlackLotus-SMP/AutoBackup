package endpoints

import (
	"backup/cfg"
	"backup/routes/controller"
	"backup/rsync"
	"github.com/gin-gonic/gin"
)

type Backup struct {
	Conf          *cfg.Config
	RSyncExecutor *rsync.Executor
}

func (rl Backup) Route(engine *gin.Engine) {
	backupController := controller.Backup{
		Conf:          rl.Conf,
		RSyncExecutor: rl.RSyncExecutor,
	}
	engine.GET("/reload", backupController.Reload)
	backupRouter := engine.Group("/backup")
	backupRouter.GET("/create/:name", backupController.Create)
}
