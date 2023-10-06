package routes

import (
	"aplikasi_penjualan_spareparts/config"
	"aplikasi_penjualan_spareparts/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {

	// 	order := []models.Order{}

	// 	config.DB.Preload("Customer").Find(&order)

	// 	getOrderResponse := []models.OrderResponse{}

	// 	for _, od := range order {
	// 		cust := models.CustomerResponse{
	// 			ID:            od.Customer.ID,
	// 			CustomerName:  od.Customer.CustomerName,
	// 			CustomerEmail: od.Customer.CustomerEmail,
	// 		}

	// 		odr := models.OrderRequest{
	// 			ID:            od.ID,
	// 			OrderDate:     od.OrderDate,
	// 			PaymentCode:   od.Paymentinfo.PaymentCode,
	// 			PaymentMethod: od.Paymentinfo.PaymentMethod,
	// 			PaymentDate:   od.Paymentinfo.PaymentDate,
	// 			Customer:      cust,
	// 		}

	// 		getOrderResponse = append(getOrderResponse, odr)
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Successfully retrieved product",
	// 		"data":    getOrderResponse,
	// 	})
	// }

	orders := []models.Order{}
	// config.DB.Find(&customers)

	// dengan relational db
	config.DB.Preload("Customer").Find(&orders)

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully retrieved orders",
		"data":    orders,
	})
}

func GetOrderById(c *gin.Context) {
	id := c.Param("id")

	order := []models.Order{}

	// dengan relational db
	data := config.DB.Preload("Customer").First(&order, "id = ?", id)

	if data.Error != nil {
		// log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "order not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully retrieved order",
		"data":    order,
	})
}

func PostOrders(c *gin.Context) {

	// ambil data post dari JSON
	var reqOrd models.OrderRequest
	c.BindJSON(&reqOrd)

	Order := models.Order{
		OrderCode:   reqOrd.OrderCode,
		OrderDate:   reqOrd.OrderDate,
		OrderStatus: reqOrd.OrderStatus,
		Qty:         reqOrd.Qty,
		CustomerID:  reqOrd.CustomerID,
		Paymentinfo: models.Paymentinfo{
			PaymentCode:   reqOrd.PaymentCode,
			PaymentMethod: reqOrd.PaymentMethod,
			PaymentDate:   reqOrd.PaymentDate,
		},
	}

	// insert data to DB
	config.DB.Create(&Order)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully created Order",
		"data":    Order,
	})

	// // ambil data post dari JSON
	// var order models.Order
	// c.BindJSON(&order)

	// // insert data to DB
	// config.DB.Create(&order)

	// // JSON response
	// c.JSON(http.StatusCreated, gin.H{
	// 	"data":    order,
	// 	"message": "Successfully created order",
	// })
}

func PutOrders(c *gin.Context) {
	id := c.Param("id")

	var order models.Order
	// data := config.DB.Where("id = ?", id).Find(&department)
	data := config.DB.First(&order, "id = ?", id)

	if data.Error != nil {
		// log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "order not found",
		})
		return
	}
	c.BindJSON(&order)

	config.DB.Model(&order).Where("id = ?", id).Updates(&order)

	c.JSON(http.StatusOK, gin.H{
		"message": "Update Successfully",
		"data":    order,
	})
}

func DeleteOrders(c *gin.Context) {
	id := c.Param("id")

	var order models.Order

	data := config.DB.First(&order, "id = ?", id)

	if data.Error != nil {
		// log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "order not found",
		})
		return
	}

	config.DB.Delete(&order, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Successfully",
	})
}
