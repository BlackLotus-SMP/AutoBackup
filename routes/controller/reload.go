package controller

import (
	"backup/routes/services"
	"github.com/gin-gonic/gin"
)

func (controller Backup) Reload(c *gin.Context) {
	res := services.BackupReload.Reload()
	c.AbortWithStatusJSON(res.Code, res)
}
