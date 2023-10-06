package models

import "gorm.io/gorm"

type City struct {
	gorm.Model
	Citycode  string `json:"citycode"`
	Cityname  string `json:"cityname"`
	Customers []Customer
}

type GetCityResponse struct {
	ID        uint   `json:"cityid"`
	Citycode  string `json:"citycode"`
	Cityname  string `json:"cityname"`
	Customers []CustomerResponse
}
