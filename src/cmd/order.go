package cmd

import (
	"PicusFinalCase/src/handler"
	"PicusFinalCase/src/pkg/graceful"
	"PicusFinalCase/src/pkg/middleware"
	"PicusFinalCase/src/repository"
	"PicusFinalCase/src/service"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

var orderCmd = &cobra.Command{
	Use:   "order",
	Short: "order api",
	Long:  `basic order api`,
	Run: func(cmd *cobra.Command, args []string) {

		cfg := LoadConfig()
		orderCfg := cfg.DomainConfigs["order"]

		db := NewSqlDb(&orderCfg.DBConfig)

		SetLog()

		gin.SetMode(gin.ReleaseMode)
		r := gin.New()

		SwaggerSettings(orderCfg.ServerConfig)
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

		srv := &http.Server{
			Addr:         fmt.Sprintf("127.0.0.1:%s", orderCfg.ServerConfig.Port),
			Handler:      r,
			ReadTimeout:  time.Duration(orderCfg.ServerConfig.ReadTimeoutSecs * int64(time.Second)),
			WriteTimeout: time.Duration(orderCfg.ServerConfig.WriteTimeoutSecs * int64(time.Second)),
		}

		rootRouter := r.Group(orderCfg.ServerConfig.RoutePrefix)
		rootRouter.Use(middleware.Recovery())
		rootRouter.Use(middleware.LoggingMiddleware())

		orderRouter := rootRouter.Group("/order")
		orderRepo := repository.NewOrderRepository(db)
		orderService := service.NewOrderService(orderRepo, productRepo, cartRepo)
		handler.NewOrderHandler(orderRouter, cfg.JWTConfig, orderService)

		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()

		graceful.ShutdownGin(srv, time.Duration(orderCfg.ServerConfig.TimeoutSecs*int64(time.Second)))

	},
}
