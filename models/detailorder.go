package models

import "gorm.io/gorm"

type Detailorder struct {
	gorm.Model
	OrderID    uint    `json:"order_id"`
	Order      Order   `gorm:"foreignKey:OrderID;references:ID"`
	ProductID  uint    `json:"product_id"`
	Product    Product `gorm:"foreignKey:ProductID;references:ID"`
	Qty        int     `json:"qty"`
	TotalPrice float64 `json:"totalprice"`
	Status     string  `json:"status"`
}

type RequestDetailOrder struct {
	OrderID    uint    `json:"order_id"`
	ProductID  uint    `json:"product_id"`
	Qty        int     `json:"qty"`
	TotalPrice float64 `json:"totalprice"`
	Status     string  `json:"status"`
}

type ResponseGetDetailOrder struct {
	ID          uint    `json:"id"`
	OrderID     uint    `json:"order_id"`
	ProductName string  `json:"product_name"`
	Qty         int     `json:"qty"`
	TotalPrice  float64 `json:"totalprice"`
	Status      string  `json:"status"`
}
