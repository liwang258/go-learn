package db

import (
	"context"
	"errors"
	"fmt"
	"go-learn/task4"
	"strings"

	"gorm.io/gorm"
)

// UserOp 这里定义了一个别名
type UserOp struct{}

// Create 保存用户信息
// dst 用户请求参数
func (u *UserOp) Create(dst interface{}) (task4.User, error) {
	user, ok := dst.(*task4.User)
	if !ok {
		return task4.User{}, errors.New("dst is not user")
	}
	if strings.TrimSpace(user.Email) == "" || strings.TrimSpace(user.Password) == "" {
		return task4.User{}, errors.New("email or password is empty")
	}
	//这里会校验表是否存在，不存在则创建表
	u.CreateTable(task4.User{})
	db := GetDb()
	//写入记录，会直接将主键设置到user.id中
	db.Create(user)
	return *user, nil
}

// Update 更新订单
func (u *UserOp) Update(dst interface{}, where interface{}) (bool, error) {
	u.CreateTable(task4.User{})
	newPwd := dst.(string)
	userId := where.(int64)

	db := GetDb()
	if _, err := gorm.G[task4.User](db).Where("id=?", userId).Update(context.Background(), "password", &newPwd); err != nil {
		return false, err
	}
	_, err := gorm.G[task4.User](db).Where("id=?", userId).First(context.Background())
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *UserOp) Delete(where interface{}) (int, error) {
	u.CreateTable(task4.User{})
	userId := where.(int64)
	db := GetDb()
	rowAffected, err := gorm.G[task4.User](db).Where(userId).Delete(db.Statement.Context)
	return int(rowAffected), err
}

func (u *UserOp) CreateTable(dst interface{}) {
	db := GetDb()
	//检测表是否存在，不存在则先创建表
	exists := db.Migrator().HasTable(dst)
	if !exists {
		fmt.Println("开始创建表")
		//创建表
		if err := db.Migrator().CreateTable(dst); err != nil {
			fmt.Printf("create table failed:%v", err)
		} else {
			fmt.Println("创建表成功")
		}
	}
}
