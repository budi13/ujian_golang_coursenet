package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	CustomerCode  string  `json:"customercode"`
	CustomerName  string  `json:"customername"`
	Address       string  `json:"address"`
	PhoneNumber   string  `json:"phonenumber"`
	CustomerEmail string  `json:"customeremail"`
	Gender        string  `json:"gender"`
	BirthOfDate   string  `json:"birthofdate"`
	CityID        uint    `json:"cityid"`
	Orders        []Order `json:"orders"`
}

type GetCustomerResponse struct {
	ID            uint            `json:"customerid"`
	CustomerCode  string          `json:"customercode"`
	CustomerName  string          `json:"customername"`
	Address       string          `json:"address"`
	PhoneNumber   string          `json:"Phonenumber"`
	CustomerEmail string          `json:"customeremail"`
	Gender        string          `json:"gender"`
	BirthOfDate   string          `json:"birthofdate"`
	CityID        uint            `json:"cityid"`
	Orders        []OrderResponse `json:"orders"`
}

type CustomerResponse struct {
	ID            uint   `json:"customerid"`
	CustomerCode  string `json:"customercode"`
	CustomerName  string `json:"customername"`
	Address       string `json:"address"`
	PhoneNumber   string `json:"Phonenumber"`
	CustomerEmail string `json:"customeremail"`
	Gender        string `json:"gender"`
	BirthOfDate   string `json:"birthofdate"`
}
