package main

import (
	"example/web-service-gin/database"
	"example/web-service-gin/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()
	database.InitDB()
	defer database.CloseDB()

	// Đăng ký các API
	r.POST("/register", handlers.RegisterUser)
	r.POST("/login", handlers.LoginUser)
	r.GET("/notifications", handlers.GetAllUserNotifications)
	r.GET("/user/:user_id/notifications", handlers.GetUserNotifications)

	// Chạy server
	r.Run(":8080")
}
