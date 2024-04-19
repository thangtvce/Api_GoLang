package handlers

import (
	"example/web-service-gin/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Password string `json:"password"`
}
type Notification struct {
	ID      int    `json:"notification_id"`
	Message string `json:"message"`
}

type Role struct {
	ID   int    `json:"role_id"`
	Name string `json:"role_name"`
}

type UserRole struct {
	UserID int `json:"user_id"`
	RoleID int `json:"role_id"`
}

type UserNotification struct {
	UserID         int    `json:"user_id"`
	NotificationID int    `json:"notification_id"`
	Message        string `json:"message"`
}

func RegisterUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := database.DB.Exec("INSERT INTO users (username, role, password) VALUES ($1, $2, $3)", user.Username, user.Role, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
func LoginUser(c *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Giải mã thông tin người dùng từ yêu cầu
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Truy vấn cơ sở dữ liệu để lấy thông tin người dùng dựa trên username
	var dbUser User
	err := database.DB.QueryRow("SELECT username, password FROM users WHERE username = $1", user.Username).Scan(&dbUser.Username, &dbUser.Password)
	if err != nil {
		// Xử lý lỗi không tìm thấy người dùng hoặc lỗi truy vấn cơ sở dữ liệu
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Kiểm tra xem password nhập vào có khớp với password trong cơ sở dữ liệu không
	if user.Password != dbUser.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Nếu xác thực thành công, bạn có thể thực hiện các hành động tiếp theo như tạo mã thông báo JWT và trả về nó cho người dùng.
	// Ví dụ:
	// token, err := generateJWTToken(user.Username)
	// if err != nil {
	//     c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
	//     return
	// }

	// Trả về thông báo thành công nếu xác thực thành công
	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully"})
}

// Lấy danh sách toàn bộ thông báo của toàn bộ người dùng (chỉ dành cho admin)
func GetAllUserNotifications(c *gin.Context) {
	// Thực hiện truy vấn để lấy danh sách toàn bộ thông báo của toàn bộ người dùng
	rows, err := database.DB.Query("SELECT un.user_id, n.notification_id, n.message FROM UserNotifications un JOIN Notifications n ON un.notification_id = n.notification_id")
	if err != nil {
		log.Println("Error querying user notifications:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user notifications"})
		return
	}
	defer rows.Close()

	// Tạo một map để lưu trữ thông báo của từng người dùng
	userNotifications := make(map[int][]Notification)

	// Lặp qua các hàng và thêm thông báo vào map tương ứng với user_id
	for rows.Next() {
		var userNotification UserNotification
		if err := rows.Scan(&userNotification.UserID, &userNotification.NotificationID, &userNotification.Message); err != nil {
			log.Println("Error scanning user notification:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan user notifications"})
			return
		}

		// Chuyển đổi từ UserNotification sang Notification
		notification := Notification{
			ID:      userNotification.NotificationID,
			Message: userNotification.Message,
		}

		// Thêm thông báo vào map tương ứng với user_id
		userNotifications[userNotification.UserID] = append(userNotifications[userNotification.UserID], notification)
	}

	// Trả về danh sách thông báo của toàn bộ người dùng
	c.JSON(http.StatusOK, userNotifications)
}

// Lấy danh sách thông báo của một người dùng cụ thể
func GetUserNotifications(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		log.Println("User ID is empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}
	rows, err := database.DB.Query("SELECT n.notification_id, n.message FROM UserNotifications un JOIN Notifications n ON un.notification_id = n.notification_id WHERE un.user_id = $1", userID)
	if err != nil {
		log.Println("Error querying user notifications:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user notifications"})
		return
	}
	defer rows.Close()

	var userNotifications []Notification

	for rows.Next() {
		var notification Notification
		if err := rows.Scan(&notification.ID, &notification.Message); err != nil {
			log.Println("Error scanning user notification:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan user notifications"})
			return
		}
		userNotifications = append(userNotifications, notification)
	}

	c.JSON(http.StatusOK, userNotifications)
}
