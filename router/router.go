package router

import (
	"github.com/edfan0930/rhino/domain"
	"github.com/gin-gonic/gin"
)

func InitRouter() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/daily/message/:text", domain.DailyMessage)

	r.Run(":8123")
}
