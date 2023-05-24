package api

import (
	"bicycle/bicycle_api_gateway/api/handler"
	"bicycle/bicycle_api_gateway/config"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(h handler.Handler, cfg config.Config) (r *gin.Engine) {
	r = gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/user")
	r.POST("/", h.CreateUser)

	return
}
