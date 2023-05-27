package routers

import (
	"Project2-Hacktiv8/controllers"
	"Project2-Hacktiv8/db"
	"github.com/gin-gonic/gin"
)

func init() {
	db.InitializeDB()
}

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders/", controllers.GetOrder)
	router.PUT("/orders/:orderId", controllers.UpdateOrder)
	router.DELETE("/orders/:orderId", controllers.DeleteOrder)

	return router
}
