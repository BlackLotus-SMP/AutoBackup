package controller

import (
	"backup/routes/services"
	"github.com/gin-gonic/gin"
)

func (controller Backup) List(c *gin.Context) {
	res := services.BackupCreate.List(controller.Conf)
	c.IndentedJSON(res.Code, res)
}
