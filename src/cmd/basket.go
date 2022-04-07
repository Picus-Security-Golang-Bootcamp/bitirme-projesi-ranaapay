package cmd

import (
	"github.com/ranaapay/bitirme-projesi-ranaapay/src/pkg/config"
	"github.com/ranaapay/bitirme-projesi-ranaapay/src/pkg/db"
	"log"
)

var (
	ConfigFile = "./pkg/config/config"
)
func Execute() {
	cfg, err := config.LoadConfig(ConfigFile)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	_, err = db.NewPsqlDB(cfg)
	if err != nil {
		log.Fatalf("DatabaseConnect: %v", err)
	}
}
