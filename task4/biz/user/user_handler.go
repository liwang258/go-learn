package user

import (
	"fmt"
	"go-learn/task4"
	"go-learn/task4/db"
)

func CreateUser(query task4.UserQuery) (task4.UserQueryResult, error) {
	user := &task4.User{
		Username: query.UserName,
		Password: query.Password,
		Email:    query.Email,
	}
	fmt.Println(user)
	var op db.BaseOp[task4.User] = &db.UserOp{}
	var u interface{} = user
	createUser, err := op.Create(u)
	if err != nil {
		return task4.UserQueryResult{}, err
	}

	return task4.UserQueryResult{
		UserId:    int64(createUser.ID),
		IsSuccess: true,
	}, nil

}
