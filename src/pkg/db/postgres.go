package db

import (
	"PicusFinalCase/src/pkg/config"
	"PicusFinalCase/src/pkg/errorHandler"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPsqlDB(cfg *config.Config) (*gorm.DB, interface{}) {

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		cfg.DBConfig.Host,
		cfg.DBConfig.Port,
		cfg.DBConfig.Username,
		cfg.DBConfig.Name,
		cfg.DBConfig.Password,
	)

	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Error("The file with the given filename could not be found at the specified address. : %v", err)
		return nil, errorHandler.GormOpenError
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Error("Failed to get sql db. : %s", err.Error())
		return nil, errorHandler.SqlDBError
	}

	if err = sqlDB.Ping(); err != nil {
		log.Error("Database connection does not exist. Ping Error : %s", err.Error())
		return nil, errorHandler.SqlDBPingError
	}
	return db, nil
}
