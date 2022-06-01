package cmd

import (
	"PicusFinalCase/src/docs"
	"PicusFinalCase/src/handler"
	"PicusFinalCase/src/pkg/config"
	"PicusFinalCase/src/pkg/db"
	"PicusFinalCase/src/pkg/graceful"
	"PicusFinalCase/src/pkg/middleware"
	"PicusFinalCase/src/repository"
	"PicusFinalCase/src/service"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	swaggerfiles "github.com/swaggo/files" // swagger embed files
	"github.com/swaggo/gin-swagger"        // gin-swagger middleware
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
)

var (
	ConfigFile = "./pkg/config/config"
)

var rootCmd = &cobra.Command{
	Use:   "basket",
	Short: "basket",
	Long:  `basic basket api`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := LoadConfig()

		db := NewSqlDb(cfg)

		SetLog()

		gin.SetMode(gin.ReleaseMode)
		r := gin.New()

		SwaggerSettings(cfg.ServerConfig)
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

		srv := &http.Server{
			Addr:         fmt.Sprintf("127.0.0.1:%s", cfg.ServerConfig.Port),
			Handler:      r,
			ReadTimeout:  time.Duration(cfg.ServerConfig.ReadTimeoutSecs * int64(time.Second)),
			WriteTimeout: time.Duration(cfg.ServerConfig.WriteTimeoutSecs * int64(time.Second)),
		}

		rootRouter := r.Group(cfg.ServerConfig.RoutePrefix)
		rootRouter.Use(middleware.Recovery())
		rootRouter.Use(middleware.LoggingMiddleware())

		authenticationRouter := rootRouter.Group("/authentication")
		authRepo := repository.NewAuthRepository(db)
		authService := service.NewAuthService(cfg.JWTConfig, authRepo)
		handler.NewAuthHandler(authenticationRouter, authService)

		categoryRouter := rootRouter.Group("/category")
		categoryRepo := repository.NewCategoryRepository(db)
		categoryService := service.NewCategoryService(&categoryRepo)
		handler.NewCategoryHandler(categoryRouter, cfg.JWTConfig, categoryService)

		productRouter := rootRouter.Group("/product")
		productRepo := repository.NewProductRepository(db)
		productService := service.NewProductService(productRepo, &categoryRepo)
		handler.NewProductHandler(productRouter, cfg.JWTConfig, productService)

		cartRouter := rootRouter.Group("/cart")
		cartRepo := repository.NewCartRepository(db)
		cartService := service.NewCartService(cartRepo, productRepo)
		handler.NewCartHandler(cartRouter, cfg.JWTConfig, cartService)

		orderRouter := rootRouter.Group("/order")
		orderRepo := repository.NewOrderRepository(db)
		orderService := service.NewOrderService(orderRepo, productRepo, cartRepo)
		handler.NewOrderHandler(orderRouter, cfg.JWTConfig, orderService)

		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()

		graceful.ShutdownGin(srv, time.Duration(cfg.ServerConfig.TimeoutSecs*int64(time.Second)))

	},
}

/*
func Execute() {

	cfg := LoadConfig()

	db := NewSqlDb(cfg)

	SetLog()

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	SwaggerSettings(cfg.ServerConfig)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	srv := &http.Server{
		Addr:         fmt.Sprintf("127.0.0.1:%s", cfg.ServerConfig.Port),
		Handler:      r,
		ReadTimeout:  time.Duration(cfg.ServerConfig.ReadTimeoutSecs * int64(time.Second)),
		WriteTimeout: time.Duration(cfg.ServerConfig.WriteTimeoutSecs * int64(time.Second)),
	}

	rootRouter := r.Group(cfg.ServerConfig.RoutePrefix)
	rootRouter.Use(middleware.Recovery())
	rootRouter.Use(middleware.LoggingMiddleware())

	authenticationRouter := rootRouter.Group("/authentication")
	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(cfg.JWTConfig, authRepo)
	handler.NewAuthHandler(authenticationRouter, authService)

	categoryRouter := rootRouter.Group("/category")
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(&categoryRepo)
	handler.NewCategoryHandler(categoryRouter, cfg.JWTConfig, categoryService)

	productRouter := rootRouter.Group("/product")
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo, &categoryRepo)
	handler.NewProductHandler(productRouter, cfg.JWTConfig, productService)

	cartRouter := rootRouter.Group("/cart")
	cartRepo := repository.NewCartRepository(db)
	cartService := service.NewCartService(cartRepo, productRepo)
	handler.NewCartHandler(cartRouter, cfg.JWTConfig, cartService)

	orderRouter := rootRouter.Group("/order")
	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo, productRepo, cartRepo)
	handler.NewOrderHandler(orderRouter, cfg.JWTConfig, orderService)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	graceful.ShutdownGin(srv, time.Duration(cfg.ServerConfig.TimeoutSecs*int64(time.Second)))
}
*/

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

// LoadConfig Read Config yaml
func LoadConfig() *config.Config {
	cfg, err := config.LoadConfig(ConfigFile)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	return cfg
}

// NewSqlDb Initialize db
func NewSqlDb(cfg *config.Config) *gorm.DB {
	db, err := db.NewPsqlDB(cfg)
	if err != nil {
		log.Fatalf("DatabaseConnect: %v", err)
	}
	return db
}

// SetLog Setup logger
func SetLog() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})
}

func SwaggerSettings(cfg config.ServerConfig) {
	docs.SwaggerInfo.Title = "Basket Application"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Description = "Basket Service"
	docs.SwaggerInfo.BasePath = cfg.RoutePrefix
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Port)
	docs.SwaggerInfo.Schemes = []string{"http"}
}
