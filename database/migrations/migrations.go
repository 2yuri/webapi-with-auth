package migrations

import (
	"github.com/hyperyuri/webapi-with-go/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
