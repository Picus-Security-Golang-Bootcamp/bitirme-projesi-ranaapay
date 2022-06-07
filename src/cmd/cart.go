package cmd

import (
	"PicusFinalCase/src/client"
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

var cartCmd = &cobra.Command{
	Use:   "cart",
	Short: "cart api",
	Long:  `basic cart api`,
	Run: func(cmd *cobra.Command, args []string) {

		cfg := LoadConfig()
		cartCfg := cfg.DomainConfigs["cart"]

		db := NewSqlDb(&cartCfg.DBConfig)

		SetLog()

		gin.SetMode(gin.ReleaseMode)
		r := gin.New()

		SwaggerSettings(cartCfg.ServerConfig)
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

		srv := &http.Server{
			Addr:         fmt.Sprintf("127.0.0.1:%s", cartCfg.ServerConfig.Port),
			Handler:      r,
			ReadTimeout:  time.Duration(cartCfg.ServerConfig.ReadTimeoutSecs * int64(time.Second)),
			WriteTimeout: time.Duration(cartCfg.ServerConfig.WriteTimeoutSecs * int64(time.Second)),
		}

		rootRouter := r.Group(cartCfg.ServerConfig.RoutePrefix)
		rootRouter.Use(middleware.Recovery())
		rootRouter.Use(middleware.LoggingMiddleware())

		cartRouter := rootRouter.Group("/cart")

		cartRepo := repository.NewCartRepository(db)

		productClient := client.NewProductClient(cfg.DomainConfigs["product"].ServerConfig, "/product")

		cartService := service.NewCartService(cartRepo, productClient)

		handler.NewCartHandler(cartRouter, cfg.JWTConfig, cartService)

		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()

		graceful.ShutdownGin(srv, time.Duration(cartCfg.ServerConfig.TimeoutSecs*int64(time.Second)))

	},
}
