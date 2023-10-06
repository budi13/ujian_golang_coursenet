package routes

import (
	"aplikasi_penjualan_spareparts/config"
	"aplikasi_penjualan_spareparts/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	category := []models.Category{}

	// dengan relational db
	config.DB.Preload("Products").Find(&category)

	getCategoryResponse := []models.GetCategoryResponse{}

	for _, cat := range category {

		products := []models.ProductResponse{}
		for _, p := range cat.Products {
			prod := models.ProductResponse{
				ID:          p.ID,
				ProductCode: p.ProductCode,
				ProductName: p.ProductName,
				Description: p.Description,
				Price:       p.Price,
				Stock:       p.Stock,
				Weight:      p.Weight,
				DateIn:      p.DateIn,
				Picture:     p.Picture,
				Purchase:    p.Purchase,
				Discount:    p.Discount,
			}

			products = append(products, prod)
		}

		ctg := models.GetCategoryResponse{
			ID:           cat.ID,
			CategoryCode: cat.CategoryCode,
			CategoryName: cat.CategoryName,
			Products:     products,
		}

		getCategoryResponse = append(getCategoryResponse, ctg)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully retrieved category",
		"data":    getCategoryResponse,
	})
}

func GetCategoriesById(c *gin.Context) {

	id := c.Param("id")

	var category models.Category

	// dengan relational db
	data := config.DB.Preload("Products").First(&category, "id = ?", id)

	if data.Error != nil {
		// log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "department not found",
		})
		return
	}

	products := []models.ProductResponse{}
	for _, p := range category.Products {
		prod := models.ProductResponse{
			ID:          p.ID,
			ProductCode: p.ProductCode,
			ProductName: p.ProductName,
			Description: p.Description,
			Price:       p.Price,
			Stock:       p.Stock,
			Weight:      p.Weight,
			DateIn:      p.DateIn,
			Picture:     p.Picture,
			Purchase:    p.Purchase,
			Discount:    p.Discount,
		}
		products = append(products, prod)
	}

	getCategoryResponse := models.GetCategoryResponse{
		ID:           category.ID,
		CategoryCode: category.CategoryCode,
		CategoryName: category.CategoryName,
		Products:     products,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully retrieved category",
		"data":    getCategoryResponse,
	})
}

func PostCategories(c *gin.Context) {

	// ambil data post dari JSON
	var category models.Category
	c.BindJSON(&category)

	// insert data to DB
	config.DB.Create(&category)

	// JSON response
	c.JSON(http.StatusCreated, gin.H{
		"data":    category,
		"message": "Successfully created category",
	})
}

func PutCategories(c *gin.Context) {
	id := c.Param("id")

	var category models.Category

	data := config.DB.First(&category, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "category not found",
		})
		return
	}
	c.BindJSON(&category)

	config.DB.Model(&category).Where("id = ?", id).Updates(&category)

	c.JSON(http.StatusOK, gin.H{
		"data":    category,
		"message": "Update Successfully",
	})
}

func DeleteCategories(c *gin.Context) {
	id := c.Param("id")

	var category models.Category

	data := config.DB.First(&category, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "category not found",
		})
		return
	}

	config.DB.Delete(&category, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Successfully",
	})
}
