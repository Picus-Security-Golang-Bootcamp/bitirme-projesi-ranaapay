package db

import (
	"PicusFinalCase/src/pkg/config"
	"PicusFinalCase/src/pkg/errorHandler"
	"fmt"
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
		return nil, errorHandler.GormOpenError
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, errorHandler.SqlDBError
	}
	if err = sqlDB.Ping(); err != nil {
		return nil, errorHandler.SqlDBPingError
	}
	return db, nil
}
