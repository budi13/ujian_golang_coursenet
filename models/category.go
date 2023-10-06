package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryCode string    `json:"categorycode"`
	CategoryName string    `json:"categoryname"`
	Products     []Product `json:"product"`
}

type CategoryResponse struct {
	ID           uint   `json:"id"`
	CategoryCode string `json:"categorycode"`
	CategoryName string `json:"categoryname"`
}

type GetCategoryResponse struct {
	ID           uint              `json:"id"`
	CategoryCode string            `json:"categorycode"`
	CategoryName string            `json:"categoryname"`
	Products     []ProductResponse `json:"product"`
}
