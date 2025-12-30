package db

import (
	"go-learn/task4"
	"strings"

	"gorm.io/gorm"
)

type user task4.User

func (u *user) Create(user *task4.User) bool {
	if strings.TrimSpace(user.Email) == "" || strings.TrimSpace(user.Password) == "" {
		return false
	}
	db := GetDb()
	gorm.G[task4.User](db).Create(db.Statement.Context, user)
	return true
}

func (u *user) Update(newPwd *string, userId *int64) {
	db := GetDb()
	gorm.G[task4.User](db).Where(userId).Update(db.Statement.Context, "password", newPwd)
}

func (u *user) Delete(userId *int64) {
	db := GetDb()
	gorm.G[task4.User](db).Where(userId).Delete(db.Statement.Context)
}
