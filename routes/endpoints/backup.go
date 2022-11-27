package endpoints

import (
	"backup/cfg"
	"backup/routes/controller"
	"github.com/gin-gonic/gin"
)

type Backup struct {
	Conf *cfg.Config
}

func (rl Backup) Route(engine *gin.Engine) {
	backupController := controller.Backup{
		Conf: rl.Conf,
	}
	engine.GET("/reload", backupController.Reload)
	backupRouter := engine.Group("/backup")
	backupRouter.GET("/create/:name", backupController.Create)
}
