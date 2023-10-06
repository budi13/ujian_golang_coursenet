package routes

import (
	"aplikasi_penjualan_spareparts/config"
	"aplikasi_penjualan_spareparts/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetDetailOrder(c *gin.Context) {
	Detailorder := []models.Detailorder{}

	config.DB.Preload(clause.Associations).Find(&Detailorder)

	responseGetDetailOrder := []models.ResponseGetDetailOrder{}

	for _, do := range Detailorder {
		rdo := models.ResponseGetDetailOrder{
			ID:          do.ID,
			OrderID:     do.Order.ID,
			ProductName: do.Product.ProductName,
			Qty:         do.Order.Qty,
			TotalPrice:  do.TotalPrice,
		}
		responseGetDetailOrder = append(responseGetDetailOrder, rdo)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Data Detail Order",
		"data":    responseGetDetailOrder,
	})
}

func DetailOrderByEmployeeId(c *gin.Context) {
	var reqDetailOrder models.RequestDetailOrder

	if err := c.ShouldBindJSON(&reqDetailOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"data":    err.Error(),
		})

		c.Abort()
		return
	}

	detorder := models.Detailorder{
		OrderID:    reqDetailOrder.OrderID,
		ProductID:  reqDetailOrder.ProductID,
		Qty:        reqDetailOrder.Qty,
		TotalPrice: reqDetailOrder.TotalPrice,
		Status:     reqDetailOrder.Status,
	}

	insert := config.DB.Create(&detorder)
	if insert.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"error":   insert.Error.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":    detorder,
		"message": "Insert Sucessfully",
	})
}
