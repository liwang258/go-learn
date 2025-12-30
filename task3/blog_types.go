package task3

import (
	"time"
)

type Base struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	Base
	Name      string    `gorm:"column:name;type:varchar(30);not null"`
	PostCnt   int       `gorm:"column:post_cnt;default:0;not null"`
	Password  string    `gorm:"column:password;default:'';not null"`
	Email     string    `gorm:"column:email;default:'';not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:null"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;default:null"`
}

type Post struct {
	Base
	Id           int    `gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	UserId       int    `grom:"column:u_id;type:bigint;not null;index:u_id;many2many:users"`
	Content      string `gorm:"column:content;type:varchar(1024);default:'';not null"`
	CommentState int    `gorm:"column:comment_status;default:0;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Comment struct {
	Base
	PostId    int    `gorm:"column:post_id;type:bigint;not null;index"`
	UserId    int    `gorm:"column:u_id;not null;type:bigint"`
	Comment   string `gorm:"column:comment;type:varchar(1024);default:''"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRalatePost struct {
	Posts    []Post
	Comments []Comment
}

type HotPost struct {
	PostId       int `gorm:"column:post_id"`
	CommentTotal int `gorm:"column:count"`
}
