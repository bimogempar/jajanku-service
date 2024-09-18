package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config *Config) (*gorm.DB, error) {
	dsn := "host=" + config.DBHost +
		" user=" + config.DBUser +
		" password=" + config.DBPassword +
		" dbname=" + config.DBName +
		" port=" + config.DBPort +
		" sslmode=disable" +
		" search_path=" + config.DBSchema
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
