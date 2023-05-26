package main

import (
	"bicycle/bicycle_api_gateway/api"
	"bicycle/bicycle_api_gateway/api/handler"
	"bicycle/bicycle_api_gateway/config"
	"bicycle/bicycle_api_gateway/grpc/client"
	"bicycle/bicycle_api_gateway/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	var loggerLevel string

	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer func() {
		err := logger.Cleanup(log)
		if err != nil {
			log.Error("!!!Cleanup->Error-->>>", logger.Error(err))
		}
	}()

	svcs, err := client.NewGrpcClient(cfg)
	if err != nil {
		log.Panic("client.NewGrpcClients", logger.Error(err))
	}

	h := handler.NewHandler(cfg, log, svcs)

	r := api.SetUpRouter(h, cfg)

	log.Info("HTTP: Server being started...", logger.String("port", cfg.ServicePort))

	err = r.Run(cfg.ServicePort)
	if err != nil {
		panic(err)
	}
}
