package main

import (
	"go-learn/task2"
)

func main() {
	// c := task2.Circle{
	// 	Radius: 1,
	// }
	// r := task2.Rectangle{
	// 	Width:  1,
	// 	Length: 2,
	// }
	// fmt.Printf("Circle area: %f,Perimeter:%f \n", c.Area(), c.Perimeter())
	// fmt.Printf("Rectangle area:%f,Perimeter:%f \n", r.Area(), r.Perimeter())

	// e := task2.Employee{
	// 	Person: task2.Person{
	// 		Age:  12,
	// 		Name: "Employee",
	// 	},
	// 	EmployeeID: 123,
	// }
	// e.PrintInfo()

	// e1 := task2.Employee1{
	// 	//匿名嵌套初始化这一步和非匿名嵌套是一样的
	// 	Person: task2.Person{
	// 		Age:  12,
	// 		Name: "Employee1",
	// 	},
	// 	EmployeeID: 123,
	// }
	// e1.PrintInfo()
	// task2.ChannelWithNoBuffer()
	//task2.ChannelWithBuffer()
	task2.Lock()
	task2.NoLock()
}
