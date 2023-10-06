package config

import (
	// "golang_basic_gin/models"

	"aplikasi_penjualan_spareparts/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	var err error

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/app_penjualan_spareparts?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to Connect database")
	}

	// Migrate the table
	DB.AutoMigrate(&models.City{}, &models.Customer{}, &models.Order{},
		&models.Product{}, &models.Category{}, &models.Paymentinfo{}, models.Detailorder{}, &models.User{})
}
