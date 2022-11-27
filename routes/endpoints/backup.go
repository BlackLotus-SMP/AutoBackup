package endpoints

import (
	"backup/routes/controller"
	"github.com/gin-gonic/gin"
)

type Backup struct {
}

func (rl Backup) Route(engine *gin.Engine) {
	backupController := controller.Backup{}
	engine.GET("/reload", backupController.Reload)
	backupRouter := engine.Group("/backup")
	backupRouter.GET("/create/:name", backupController.Create)
}
