package controller

import (
	"backup/routes/services"
	"github.com/gin-gonic/gin"
)

func (controller Backup) Delete(c *gin.Context) {
	name := c.Param("name")
	res := services.BackupDelete.Delete(name, controller.Conf)
	c.IndentedJSON(res.Code, res)
}
