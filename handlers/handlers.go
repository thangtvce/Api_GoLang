package handlers

import (
	"example/web-service-gin/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type Notification struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

func RegisterUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := database.DB.Exec("INSERT INTO users (username, role) VALUES ($1, $2)", user.Username, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
func LoginUser(c *gin.Context) {
	// Xử lý đăng nhập (không cần thiết ở đây)
	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully"})
}

// Lấy danh sách toàn bộ thông báo của toàn bộ người dùng (chỉ dành cho admin)
func GetAllNotifications(c *gin.Context) {
	// Kiểm tra vai trò của người dùng
	// (Chỉ cho phép xem danh sách thông báo nếu vai trò là "admin")
	role := c.GetHeader("Role")
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}
	// Lấy danh sách thông báo từ cơ sở dữ liệu
	rows, err := database.DB.Query("SELECT id, message FROM notifications")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notifications"})
		return
	}
	defer rows.Close()
	var notifications []Notification
	for rows.Next() {
		var notification Notification
		if err := rows.Scan(&notification.ID, &notification.Message); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan notifications"})
			return
		}
		notifications = append(notifications, notification)
	}
	c.JSON(http.StatusOK, notifications)
}

// Lấy danh sách thông báo của một người dùng cụ thể
func GetUserNotifications(c *gin.Context) {
	userID := c.Param("id")
	// Lấy danh sách thông báo của người dùng từ cơ sở dữ liệu
	rows, err := database.DB.Query("SELECT n.id, n.message FROM notifications n JOIN user_notifications un ON n.id = un.notification_id WHERE un.user_id = $1", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user notifications"})
		return
	}
	defer rows.Close()
	var notifications []Notification
	for rows.Next() {
		var notification Notification
		if err := rows.Scan(&notification.ID, &notification.Message); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan user notifications"})
			return
		}
		notifications = append(notifications, notification)
	}
	c.JSON(http.StatusOK, notifications)
}
