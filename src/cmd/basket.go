package cmd

import (
	"PicusFinalCase/src/docs"
	"PicusFinalCase/src/pkg/config"
	"PicusFinalCase/src/pkg/db"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"os"
)

var (
	ConfigFile = "./pkg/config/config"
)

var rootCmd = &cobra.Command{
	Use:   "basket",
	Short: "basket",
	Long:  `basic basket api`,
}

func init() {
	rootCmd.AddCommand(
		authenticationCmd,
		categoryCmd,
		cartCmd,
		orderCmd,
		productCmd,
	)
}

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
func NewSqlDb(cfg *config.DBConfig) *gorm.DB {
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
