package migrations

import (
	"gin-api/database/models"

	"gorm.io/gorm"
)

func Up(db *gorm.DB) error {
	// Tạo bảng User
	return db.AutoMigrate(&models.User{})
}

func Down(db *gorm.DB) error {
	// Xóa bảng User
	return db.Migrator().DropTable(&models.User{})
}
