package routes

import (
	"aplikasi_penjualan_spareparts/config"
	"aplikasi_penjualan_spareparts/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCities(c *gin.Context) {
	city := []models.City{}

	// dengan relational db
	config.DB.Preload("Customers").Find(&city)

	getCityResponse := []models.GetCityResponse{}

	for _, ci := range city {

		customers := []models.CustomerResponse{}
		for _, cs := range ci.Customers {
			cust := models.CustomerResponse{
				ID:            cs.ID,
				CustomerCode:  cs.CustomerCode,
				CustomerName:  cs.CustomerName,
				Address:       cs.Address,
				PhoneNumber:   cs.PhoneNumber,
				CustomerEmail: cs.CustomerEmail,
				Gender:        cs.Gender,
				BirthOfDate:   cs.BirthOfDate,
			}

			customers = append(customers, cust)
		}

		cty := models.GetCityResponse{
			ID:        ci.ID,
			Citycode:  ci.Citycode,
			Cityname:  ci.Cityname,
			Customers: customers,
		}

		getCityResponse = append(getCityResponse, cty)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully retrieved city",
		"data":    getCityResponse,
	})
}

func GetCityById(c *gin.Context) {
	id := c.Param("id")

	cities := []models.City{}

	// dengan relational db
	data := config.DB.Preload("Customers").First(&cities, "id = ?", id)

	getCityResponse := []models.GetCityResponse{}

	for _, ci := range cities {

		customers := []models.CustomerResponse{}
		for _, cs := range ci.Customers {
			cust := models.CustomerResponse{
				ID:            cs.ID,
				CustomerCode:  cs.CustomerCode,
				CustomerName:  cs.CustomerName,
				Address:       cs.Address,
				PhoneNumber:   cs.PhoneNumber,
				CustomerEmail: cs.CustomerEmail,
				Gender:        cs.Gender,
				BirthOfDate:   cs.BirthOfDate,
			}

			customers = append(customers, cust)
		}

		cty := models.GetCityResponse{
			ID:        ci.ID,
			Citycode:  ci.Citycode,
			Cityname:  ci.Cityname,
			Customers: customers,
		}

		getCityResponse = append(getCityResponse, cty)
	}

	if data.Error != nil {
		// log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "customer not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully retrieved customer",
		"data":    getCityResponse,
	})
}

func PostCities(c *gin.Context) {

	// ambil data post dari JSON
	var city models.City
	c.BindJSON(&city)

	// insert data to DB
	config.DB.Create(&city)

	// JSON response
	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully created city",
		"data":    city,
	})
}

func PutCities(c *gin.Context) {
	id := c.Param("id")

	var city models.City

	data := config.DB.First(&city, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "city not found",
		})
		return
	}
	c.BindJSON(&city)

	config.DB.Model(&city).Where("id = ?", id).Updates(&city)

	c.JSON(http.StatusOK, gin.H{
		"message": "Update Successfully",
		"data":    city,
	})
}

func DeleteCities(c *gin.Context) {
	id := c.Param("id")

	var city models.Customer

	data := config.DB.First(&city, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "city not found",
		})
		return
	}

	config.DB.Delete(&city, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Successfully",
	})
}
