package routes

import (
	"aplikasi_penjualan_spareparts/config"
	"aplikasi_penjualan_spareparts/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	product := []models.Product{}

	config.DB.Preload("Category").Find(&product)

	getProductResponse := []models.GetProductResponse{}

	for _, cat := range product {
		category := models.CategoryResponse{
			ID:           cat.Category.ID,
			CategoryCode: cat.Category.CategoryCode,
			CategoryName: cat.Category.CategoryName,
		}

		prod := models.GetProductResponse{
			ID:          cat.ID,
			ProductCode: cat.ProductCode,
			ProductName: cat.ProductName,
			Description: cat.Description,
			Price:       cat.Price,
			Stock:       cat.Stock,
			Weight:      cat.Weight,
			DateIn:      cat.DateIn,
			Picture:     cat.Picture,
			Purchase:    cat.Purchase,
			Discount:    cat.Discount,
			Category:    category,
		}

		getProductResponse = append(getProductResponse, prod)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully retrieved product",
		"data":    getProductResponse,
	})
}

func GetProductById(c *gin.Context) {
	id := c.Param("id")

	var product models.Product

	// dengan relational db
	data := config.DB.Preload("Category").First(&product, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "product not found",
		})
		return
	}
	category := models.CategoryResponse{
		ID:           product.Category.ID,
		CategoryCode: product.Category.CategoryCode,
		CategoryName: product.Category.CategoryName,
	}

	getProductResponse := models.GetProductResponse{
		ID:          product.ID,
		ProductCode: product.ProductCode,
		ProductName: product.ProductName,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Weight:      product.Weight,
		DateIn:      product.DateIn,
		Picture:     product.Picture,
		Purchase:    product.Purchase,
		Discount:    product.Discount,
		Category:    category,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully retrieved product",
		"data":    getProductResponse,
	})
}

func PostProducts(c *gin.Context) {

	// ambil data post dari JSON
	var product models.Product
	c.BindJSON(&product)

	// insert data to DB
	config.DB.Create(&product)

	// JSON response
	c.JSON(http.StatusCreated, gin.H{
		"data":    product,
		"message": "Successfully created product",
	})
}

func PutProducts(c *gin.Context) {
	id := c.Param("id")

	var product models.Product

	data := config.DB.First(&product, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "product not found",
		})
		return
	}
	c.BindJSON(&product)

	config.DB.Model(&product).Where("id = ?", id).Updates(&product)

	c.JSON(http.StatusOK, gin.H{
		"data":    product,
		"message": "Update Successfully",
	})
}

func DeleteProducts(c *gin.Context) {
	id := c.Param("id")

	var product models.Product

	data := config.DB.First(&product, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "category not found",
		})
		return
	}

	config.DB.Delete(&product, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Successfully",
	})
}
