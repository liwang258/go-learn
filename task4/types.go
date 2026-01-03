package task4

import (
	"gorm.io/gorm"
)

type User struct {
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	gorm.Model
}

type Post struct {
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
	UserID  uint
	User    User
	gorm.Model
}

type Comment struct {
	Content string `gorm:"not null"`
	UserID  uint
	User    User
	PostID  uint
	Post    Post
	gorm.Model
}
