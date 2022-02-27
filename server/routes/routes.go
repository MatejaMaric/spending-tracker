package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/MatejaMaric/spending-tracker/server/controllers"
)

func Setup(server *gin.Engine) {
	server.GET("/api/transaction", controllers.GetTransactions)
	server.POST("/api/transaction", controllers.CreateTransaction)
}
