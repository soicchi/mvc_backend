package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/soicchi/chatapp_backend/pkg/models"
)

var db *gorm.DB

func SetupDB() (err error) {
	connectInfo := fmt.Sprintf(
		"%s:%s@tcp(db)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err = gorm.Open(mysql.Open(connectInfo), &gorm.Config{})
	if err != nil {
		return err
	}

	ExecMigration()

	fmt.Println("Database connected!")
	return nil
}

func GetDB() *gorm.DB {
	return db
}

func ExecMigration() {
	// 新たにmodelを追加したらここに記載
	db.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Room{},
	)
}
