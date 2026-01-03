package db

// BaseOp 定义泛型，即返回类型由实现来决定
type BaseOp[T any] interface {
	Create(dst interface{}) (T, error)
	Update(dst interface{}, where interface{}) (bool, error)
	Delete(where interface{}) (int, error)
	CreateTable(dst interface{})
}
