package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderCode   string      `json:"ordercode"` //gorm:"primaryKey:autoincrement"
	OrderDate   string      `json:"orderdate"`
	OrderStatus string      `json:"orderstatus"`
	Qty         int         `json:"qty"`
	CustomerID  uint        `json:"customerid"`         // gorm:"column:id_customer"`
	Customer    Customer    `json:"customer"`           // gorm:"foreignKey:CustomerID;references:ID"`
	Paymentinfo Paymentinfo `gorm:"foreignKey:OrderID"` // One-to-One relationship
	Products    []*Product  `gorm:"many2many:Detailorders;"`
}

type OrderResponse struct {
	ID          uint   `json:"id"`
	OrderCode   string `json:"ordercode"`
	OrderDate   string `json:"orderdate"`
	OrderStatus string `json:"orderstatus"`
}

type OrderRequest struct {
	ID            uint   `json:"id"`
	OrderCode     string `json:"ordercode"` //gorm:"primaryKey:autoincrement"
	OrderDate     string `json:"orderdate"`
	OrderStatus   string `json:"orderstatus"`
	Qty           int    `json:"qty"`
	CustomerID    uint   `json:"customerid"` // gorm:"column:id_customer"`
	PaymentCode   string `json:"paymentcode"`
	PaymentMethod string `json:"paymentmethod"`
	PaymentDate   string `json:"paymentdate"`
	Customer      CustomerResponse
}
