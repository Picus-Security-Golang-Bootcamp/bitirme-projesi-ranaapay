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

var authenticationCmd = &cobra.Command{
	Use:   "auth",
	Short: "auth api",
	Long:  `basic auth api`,
	Run: func(cmd *cobra.Command, args []string) {

		cfg := LoadConfig()
		authCfg := cfg.DomainConfigs["authentication"]

		db := NewSqlDb(&authCfg.DBConfig)

		SetLog()

		gin.SetMode(gin.ReleaseMode)
		r := gin.New()

		SwaggerSettings(authCfg.ServerConfig)
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

		srv := &http.Server{
			Addr:         fmt.Sprintf("127.0.0.1:%s", authCfg.ServerConfig.Port),
			Handler:      r,
			ReadTimeout:  time.Duration(authCfg.ServerConfig.ReadTimeoutSecs * int64(time.Second)),
			WriteTimeout: time.Duration(authCfg.ServerConfig.WriteTimeoutSecs * int64(time.Second)),
		}

		rootRouter := r.Group(authCfg.ServerConfig.RoutePrefix)
		rootRouter.Use(middleware.Recovery())
		rootRouter.Use(middleware.LoggingMiddleware())

		authenticationRouter := rootRouter.Group("/authentication")
		authRepo := repository.NewAuthRepository(db)
		authService := service.NewAuthService(cfg.JWTConfig, authRepo)
		handler.NewAuthHandler(authenticationRouter, authService)

		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()

		graceful.ShutdownGin(srv, time.Duration(authCfg.ServerConfig.TimeoutSecs*int64(time.Second)))

	},
}
