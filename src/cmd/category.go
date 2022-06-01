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

var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "category api",
	Long:  `basic category api`,
	Run: func(cmd *cobra.Command, args []string) {

		cfg := LoadConfig()
		categoryCfg := cfg.DomainConfigs["category"]

		db := NewSqlDb(&categoryCfg.DBConfig)

		SetLog()

		gin.SetMode(gin.ReleaseMode)
		r := gin.New()

		SwaggerSettings(categoryCfg.ServerConfig)
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

		srv := &http.Server{
			Addr:         fmt.Sprintf("127.0.0.1:%s", categoryCfg.ServerConfig.Port),
			Handler:      r,
			ReadTimeout:  time.Duration(categoryCfg.ServerConfig.ReadTimeoutSecs * int64(time.Second)),
			WriteTimeout: time.Duration(categoryCfg.ServerConfig.WriteTimeoutSecs * int64(time.Second)),
		}

		rootRouter := r.Group(categoryCfg.ServerConfig.RoutePrefix)
		rootRouter.Use(middleware.Recovery())
		rootRouter.Use(middleware.LoggingMiddleware())

		categoryRouter := rootRouter.Group("/category")
		categoryRepo := repository.NewCategoryRepository(db)
		categoryService := service.NewCategoryService(&categoryRepo)
		handler.NewCategoryHandler(categoryRouter, cfg.JWTConfig, categoryService)

		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()

		graceful.ShutdownGin(srv, time.Duration(categoryCfg.ServerConfig.TimeoutSecs*int64(time.Second)))

	},
}
