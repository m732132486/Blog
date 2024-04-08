package models

// User 注册参数
type User struct {
	UserID   int64  `gorm:"user_id"`
	Username string `gorm:"username" gorm:"unique;not null"`
	Password string `gorm:"password" gorm:"not null"`
	Token    string `gorm:"-"`
}
