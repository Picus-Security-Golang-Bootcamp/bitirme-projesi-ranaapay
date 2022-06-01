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

var productCmd = &cobra.Command{
	Use:   "product",
	Short: "product api",
	Long:  `basic product api`,
	Run: func(cmd *cobra.Command, args []string) {

		cfg := LoadConfig()
		productCfg := cfg.DomainConfigs["product"]

		db := NewSqlDb(&productCfg.DBConfig)

		SetLog()

		gin.SetMode(gin.ReleaseMode)
		r := gin.New()

		SwaggerSettings(productCfg.ServerConfig)
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

		srv := &http.Server{
			Addr:         fmt.Sprintf("127.0.0.1:%s", productCfg.ServerConfig.Port),
			Handler:      r,
			ReadTimeout:  time.Duration(productCfg.ServerConfig.ReadTimeoutSecs * int64(time.Second)),
			WriteTimeout: time.Duration(productCfg.ServerConfig.WriteTimeoutSecs * int64(time.Second)),
		}

		rootRouter := r.Group(productCfg.ServerConfig.RoutePrefix)
		rootRouter.Use(middleware.Recovery())
		rootRouter.Use(middleware.LoggingMiddleware())

		productRouter := rootRouter.Group("/product")
		productRepo := repository.NewProductRepository(db)
		productService := service.NewProductService(productRepo, &categoryRepo)
		handler.NewProductHandler(productRouter, cfg.JWTConfig, productService)

		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()

		graceful.ShutdownGin(srv, time.Duration(productCfg.ServerConfig.TimeoutSecs*int64(time.Second)))

	},
}
