package main

import (
	"aplikasi_penjualan_spareparts/config"
	"aplikasi_penjualan_spareparts/middlewares"
	"aplikasi_penjualan_spareparts/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	config.InitDB()

	r := gin.Default()

	r.GET("/home", getHome)

	api := r.Group("/api/v1")
	{
		user := api.Group("/user")
		{
			user.POST("/register", routes.RegisterUser)
			user.POST("/login", routes.GenerateToken)
		}

		customers := api.Group("/customers").Use(middlewares.Auth())
		{
			customers.GET("/", routes.GetCustomers)
			customers.GET("/:id", routes.GetCustomerById)
			customers.POST("/", routes.PostCustomers)
			customers.PUT("/:id", routes.PutCustomers)
			customers.DELETE("/:id", routes.DeleteCustomers)
		}

		orders := api.Group("orders").Use(middlewares.Auth())
		{
			orders.GET("/", routes.GetOrders)
			orders.GET("/:id", routes.GetOrderById)
			orders.POST("/", routes.PostOrders)
			orders.PUT("/:id", routes.PutOrders)
			orders.DELETE("/:id", routes.DeleteOrders)
		}

		products := api.Group("products").Use(middlewares.Auth())
		{
			products.GET("/", routes.GetProducts)
			products.GET("/:id", routes.GetProductById)
			products.POST("/", routes.PostProducts)
			products.PUT("/:id", routes.PutProducts)
			products.DELETE("/:id", routes.DeleteProducts)
		}

		categories := api.Group("categories").Use(middlewares.Auth())
		{
			categories.GET("/", routes.GetCategories)
			categories.GET("/:id", routes.GetCategoriesById)
			categories.POST("/", routes.PostCategories)
			categories.PUT("/:id", routes.PutCategories)
			categories.DELETE("/:id", routes.DeleteCategories)
		}

		cities := api.Group("cities").Use(middlewares.Auth())
		{
			cities.GET("/", routes.GetCities)
			cities.GET("/:id", routes.GetCityById)
			cities.POST("/", routes.PostCities)
			cities.PUT("/:id", routes.PutCities)
			cities.DELETE("/:id", routes.DeleteCities)
		}

		detailorder := api.Group("detailorder").Use(middlewares.Auth())
		{
			detailorder.GET("/", routes.GetDetailOrder)
			detailorder.POST("/order", routes.DetailOrderByEmployeeId)
		}
	}

	r.Run() // Listen and serve on 0.0.0.0:8080
}

func getHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome Home",
	})

}
