package task4

type UserQuery struct {
	UserId int64 `json:"userId"`
	//binding:"required"：表示该参数为必填项，否则会返回错误。
	UserName string `json:"userName" binding:"required"`
	Password string `json:"pwd"`
	Email    string `json:"email"`
}

type UserQueryResult struct {
	UserId    int64 `json:"userId"`
	IsSuccess bool  `json:"isSuccess"`
}
