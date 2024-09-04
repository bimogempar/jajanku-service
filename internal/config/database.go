package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config LoadConfig) (*gorm.DB, error) {
	dsn := "host=" + config.Database.Host + " user=" + config.Database.Username + " password=" + config.Database.Password + " dbname=" + config.Database.Name + " port=" + config.Database.Port + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
