package task3

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

func AddUser(name string) (int, error) {
	db := ConnectDb()
	user := User{
		Name: name,
	}
	CreateTable(db, &User{})
	if err := gorm.G[User](db).Create(db.Statement.Context, &user); err != nil {
		return 0, err
	} else {
		return user.Id, nil
	}
}

func AddPost(user User, content string) (int, error) {
	post := Post{
		Content: content,
		UserId:  user.Id,
	}
	db := ConnectDb()
	CreateTable(db, &Post{})
	if err := gorm.G[Post](db).Create(db.Statement.Context, &post); err != nil {
		return 0, err
	} else {
		return post.Id, nil
	}
}

func AddComment(user User, postId int, comment string) bool {
	c := Comment{
		Comment: comment,
		UserId:  user.Id,
		PostId:  postId,
	}
	db := ConnectDb()
	CreateTable(db, &Comment{})
	if err := gorm.G[Comment](db).Create(db.Statement.Context, &c); err != nil {
		return false
	} else {
		return true
	}
}

func DeleteCommet(user User, postId int) bool {
	db := ConnectDb()
	if _, err := gorm.G[Comment](db.Unscoped()).
		Where("u_id=? and post_id=?", user.Id, user.Id).
		Delete(context.Background()); err == nil {
		return true
	} else {
		return false
	}
}

func FindUser(uId int) User {
	db := ConnectDb()
	if user, err := gorm.G[User](db).Where("id=?", uId).First(db.Statement.Context); err == nil {
		return user
	} else {
		return User{}
	}

}

func FindHotPosts() *Post {
	db := ConnectDb()
	hot := HotPost{}
	//找到评论数最多的post_id
	db.Table("comments").
		Select("Count(1) as count,post_id").
		Group("post_id").
		Order("count desc").Limit(1).
		Scan(&hot)
	fmt.Printf("hot post_Id:%d,comment total:%d \n", hot.PostId, hot.CommentTotal)
	if hot.PostId > 0 {
		if p, err := gorm.G[Post](db).
			Where("id=?", hot.PostId).
			First(context.Background()); err == nil {
			fmt.Printf("post_id:%d,post_content:%s \n", hot.PostId, p.Content)
			return &p
		} else {
			fmt.Printf("post_id:%d,error:%s \n", hot.PostId, err.Error())
			return nil
		}
	} else {
		return nil
	}
}

func (post *Post) BeforeCreate(tx *gorm.DB) (err error) {
	db := tx.Statement.DB
	var affectRows int
	if affectRows, err = gorm.G[User](db).
		Where("id=?", post.UserId).
		Update(context.Background(), "post_cnt", gorm.Expr("post_cnt+?", 1)); err == nil {
		fmt.Printf("Post新增执行钩子函数，更新记录数:%d,u_id:%d \n", affectRows, post.UserId)
	}
	return err
}

func (comment *Comment) BeforeDelete(tx *gorm.DB) (err error) {
	db := tx.Statement.DB
	var result int64
	if err = tx.First(comment).Error; err != nil {
		// 若查询失败（如记录已被删除），返回错误终止删除操作
		return fmt.Errorf("failed to load comment data: %w", err)
	}
	result, err = gorm.G[Comment](db).Where("post_id=?", comment.PostId).Count(context.Background(), "post_id")
	if err != nil {
		return err
	}
	if result <= 1 {
		if affectRows, err := gorm.G[Post](db).Where("id=?", comment.PostId).
			Update(context.Background(), "comment_status", 0); err == nil {
			fmt.Printf("Comment删除执行钩子函数，更新记录数:%d,post_id:%d \n", affectRows, comment.PostId)
		}
	}
	return nil
}

func (comment *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	db := tx.Statement.DB
	var affectRows int
	if affectRows, err = gorm.G[Post](db).Where("id=?", comment.PostId).
		Update(context.Background(), "comment_status", 1); err == nil {
		fmt.Printf("Comment新增执行钩子函数，更新记录数:%d,post_id:%d \n", affectRows, comment.PostId)
	}

	return err
}
