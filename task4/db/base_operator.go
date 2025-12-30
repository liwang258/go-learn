package db

type BaseOp interface {
	Create(dst *interface{}) bool
	Update(dst *interface{}, where *[]interface{})
	Delete(dst *interface{}, where *[]interface{})
}
