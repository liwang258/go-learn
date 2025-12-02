package task2

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person     Person
	EmployeeID int
}

type Employee1 struct {
	Person
	EmployeeID int
}

func (e *Employee) PrintInfo() {
	fmt.Printf("Name:%s,Age:%d,EmployeeID:%d \n", e.Person.Name, e.Person.Age, e.EmployeeID)
}

func (e *Employee1) PrintInfo() {
	//匿名嵌套的情况下，可以直接访问被嵌套的字段，等同于e.Persion.Name
	fmt.Printf("Name:%s,Age:%d,EmployeeID:%d \n", e.Name, e.Age, e.EmployeeID)
}
