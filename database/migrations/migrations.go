package migrations

import (
	"github.com/rafamarquesrmb/webapi_books_crud_golang/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {

	db.AutoMigrate(&models.Book{})
}
