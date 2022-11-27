package controller

import (
	"backup/routes/services"
	"github.com/gin-gonic/gin"
)

func (controller Backup) Reload(c *gin.Context) {
	res := services.BackupReload.Reload(controller.Conf)
	c.IndentedJSON(res.Code, res)
}
