package models

import "gorm.io/gorm"

type Paymentinfo struct {
	gorm.Model
	PaymentCode   string `json:"paymentcode"`
	PaymentMethod string `json:"paymentmethod"`
	PaymentDate   string `json:"paymentdate"`
	OrderID       uint   // One-to-One relationship
}

type RespGetPayInfo struct {
	ID            uint   `json:"id"`
	PaymentCode   string `json:"paymentcode"`
	PaymentMethod string `json:"paymentmethod"`
	PaymentDate   string `json:"paymentdate"`
}
