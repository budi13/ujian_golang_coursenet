package routes

import (
	"aplikasi_penjualan_spareparts/config"
	"aplikasi_penjualan_spareparts/models"

	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCustomers(c *gin.Context) {

	customers := []models.Customer{}

	// dengan relational db
	config.DB.Preload("Orders").Find(&customers)

	GetCustomerResponse := []models.GetCustomerResponse{}

	for _, cs := range customers {

		orders := []models.OrderResponse{}
		for _, o := range cs.Orders {
			ord := models.OrderResponse{
				ID:          o.ID,
				OrderCode:   o.OrderCode,
				OrderDate:   o.OrderDate,
				OrderStatus: o.OrderStatus,
			}

			orders = append(orders, ord)
		}

		cust := models.GetCustomerResponse{
			ID:            cs.ID,
			CustomerCode:  cs.CustomerCode,
			CustomerName:  cs.CustomerName,
			Address:       cs.Address,
			PhoneNumber:   cs.PhoneNumber,
			CustomerEmail: cs.CustomerEmail,
			Gender:        cs.Gender,
			BirthOfDate:   cs.BirthOfDate,
			CityID:        cs.CityID,
			Orders:        orders,
		}

		GetCustomerResponse = append(GetCustomerResponse, cust)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully retrieved customers",
		"data":    GetCustomerResponse,
	})
}

func GetCustomerById(c *gin.Context) {
	id := c.Param("id")

	customers := []models.Customer{}

	// dengan relational db
	data := config.DB.Preload("Orders").First(&customers, "id = ?", id)

	GetCustomerResponse := []models.GetCustomerResponse{}

	for _, cs := range customers {

		orders := []models.OrderResponse{}
		for _, o := range cs.Orders {
			ord := models.OrderResponse{
				ID:          o.ID,
				OrderDate:   o.OrderDate,
				OrderStatus: o.OrderStatus,
			}

			orders = append(orders, ord)
		}

		cust := models.GetCustomerResponse{
			ID:            cs.ID,
			CustomerCode:  cs.CustomerCode,
			CustomerName:  cs.CustomerName,
			Address:       cs.Address,
			PhoneNumber:   cs.PhoneNumber,
			CustomerEmail: cs.CustomerEmail,
			Gender:        cs.Gender,
			BirthOfDate:   cs.BirthOfDate,
			CityID:        cs.CityID,
			Orders:        orders,
		}

		GetCustomerResponse = append(GetCustomerResponse, cust)
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
		"data":    GetCustomerResponse,
	})
}

func PostCustomers(c *gin.Context) {

	// ambil data post dari JSON
	var customer models.Customer
	c.BindJSON(&customer)

	// insert data to DB
	config.DB.Create(&customer)

	// JSON response
	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully created customer",
		"data":    customer,
	})
}

func PutCustomers(c *gin.Context) {
	id := c.Param("id")

	var customer models.Customer

	data := config.DB.First(&customer, "id = ?", id)

	if data.Error != nil {
		// log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "customer not found",
		})
		return
	}
	c.BindJSON(&customer)

	config.DB.Model(&customer).Where("id = ?", id).Updates(&customer)

	c.JSON(http.StatusOK, gin.H{
		"message": "Update Successfully",
		"data":    customer,
	})
}

func DeleteCustomers(c *gin.Context) {
	id := c.Param("id")

	var customer models.Customer

	data := config.DB.First(&customer, "id = ?", id)

	if data.Error != nil {
		// log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "customer not found",
		})
		return
	}

	config.DB.Delete(&customer, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Successfully",
	})
}
