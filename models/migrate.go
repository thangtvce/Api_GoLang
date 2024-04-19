package models

import "gorm.io/gorm"

// Khởi tạo cơ sở dữ liệu và liên kết các bảng
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Notification{}, &User{}, &NotificationUser{})
	// Các ràng buộc hoặc chỉnh sửa khác có thể được thêm vào ở đây
}
