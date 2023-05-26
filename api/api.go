package api

import (
	"bicycle/bicycle_api_gateway/api/docs"
	"bicycle/bicycle_api_gateway/api/handler"
	"bicycle/bicycle_api_gateway/config"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @description This is api gateway
func SetUpRouter(h handler.Handler, cfg config.Config) (r *gin.Engine) {

	r = gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	docs.SwaggerInfo.Title = cfg.ServiceName
	docs.SwaggerInfo.Version = cfg.Version
	// docs.SwaggerInfo.Host = cfg.ServiceHost + cfg.HTTPPort
	// docs.SwaggerInfo.Schemes = []string{cfg.HTTPScheme}

	// user api
	r.GET("/user/:id", h.GetUserById)
	r.POST("/user", h.CreateUser)
	r.DELETE("/user/:id", h.DeleteUser)

	// order api
	r.GET("order/:id", h.GetOrderById)
	r.POST("order", h.CreateOrder)
	r.DELETE("order/:id", h.DeleteOrder)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return
}
