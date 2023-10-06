package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductCode string    `json:"productcode"`
	ProductName string    `json:"productname"`
	Description string    `json:"description"`
	Price       string    `json:"price"`
	Stock       string    `json:"stock"`
	Weight      string    `json:"weight"`
	DateIn      time.Time `json:"datein"`
	Picture     string    `json:"picture"`
	Purchase    string    `json:"purchase"`
	Discount    string    `json:"discount"`
	CategoryID  uint      `json:"categoryid"`
	Category    Category  `json:"category"`
	Orders      []*Order  `gorm:"many2many:Detailorders;"`
}

type GetProductResponse struct {
	ID          uint             `json:"id"`
	ProductCode string           `json:"productcode"`
	ProductName string           `json:"productname"`
	Description string           `json:"description"`
	Price       string           `json:"price"`
	Stock       string           `json:"stock"`
	Weight      string           `json:"weight"`
	DateIn      time.Time        `json:"datein"`
	Picture     string           `json:"picture"`
	Purchase    string           `json:"purchase"`
	Discount    string           `json:"discount"`
	CategoryID  uint             `json:"category_id"`
	Category    CategoryResponse `json:"categories"`
}

type ProductResponse struct {
	ID          uint      `json:"id"`
	ProductCode string    `json:"productcode"`
	ProductName string    `json:"productname"`
	Description string    `json:"description"`
	Price       string    `json:"price"`
	Stock       string    `json:"stock"`
	Weight      string    `json:"weight"`
	DateIn      time.Time `json:"datein"`
	Picture     string    `json:"picture"`
	Purchase    string    `json:"purchase"`
	Discount    string    `json:"discount"`
}
