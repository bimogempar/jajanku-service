package internal

import "gorm.io/gorm"

func RunMigration(db *gorm.DB ,dto... interface{}){
	for _, d := range dto {
		db.AutoMigrate(d)
	}
}