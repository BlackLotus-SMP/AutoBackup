package controller

import (
	"backup/routes/services"
	"github.com/gin-gonic/gin"
)

func (controller Backup) Create(c *gin.Context) {
	name := c.Param("name")
	res := services.BackupCreate.Create(name)
	c.AbortWithStatusJSON(res.Code, res)
}
