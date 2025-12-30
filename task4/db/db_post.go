package db

import (
	"context"
	"go-learn/task4"
	"strings"

	"gorm.io/gorm"
)

type postDb task4.Post

func (p *postDb) Create(post *task4.Post) bool {
	if strings.TrimSpace(post.Content) == "" || strings.TrimSpace(post.Title) == "" {
		return false
	}
	db := GetDb()
	gorm.G[task4.Post](db).Create(db.Statement.Context, post)
	return true
}

func FindPostByPostId(postId int) (task4.Post, error) {

	db := GetDb()
	post, err := gorm.G[task4.Post](db).Where("id=?", postId).First(db.Statement.Context)
	return post, err
}

func FindPostByUserId(userId int) ([]task4.Post, error) {
	db := GetDb()
	return gorm.G[task4.Post](db).Where("user_id", userId).Find(db.Statement.Context)
}

func UpdateByPostId(postId, post task4.Post) bool {
	db := GetDb()
	if _, err := gorm.G[task4.Post](db).
		Where("id=?", postId).Updates(context.Background(), post); err == nil {
		return true
	} else {
		return false
	}
}

func DeletePostByIdAndUserId(postId int, userId int) bool {
	db := GetDb()
	if _, err := gorm.G[task4.Post](db).
		Where("post_id=? and user_id", postId, userId).
		Delete(db.Statement.Context); err == nil {
		return true
	} else {
		return false
	}

}
