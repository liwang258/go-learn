package task3

type User struct {
	Id      int    `gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Name    string `gorm:"column:name;type:varchar(30);not null"`
	PostCnt int    `gorm:"comumn:post_cnt;default:0;not null"`
}

type Post struct {
	Id           int    `gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	UserId       int    `grom:"column:u_id;type:bigint;not null;index:u_id;many2many:users"`
	Content      string `gorm:"column:content;type:varchar(1024);default:'';not null"`
	CommentState int    `gorm:"column:comment_status;default:0;not null"`
}

type Comment struct {
	Id      int    `gorm:"column:id;primaryKey;autoIncrement;type:bigint"`
	PostId  int    `gorm:"column:post_id;type:bigint;not null;index"`
	UserId  int    `gorm:"column:u_id;not null;type:bigint"`
	Comment string `gorm:"column:comment;type:varchar(1024);default:''"`
}

type UserRalatePost struct {
	Posts    []Post
	Comments []Comment
}

type HotPost struct {
	PostId       int `gorm:"column:post_id"`
	CommentTotal int `gorm:"column:count"`
}
