package task3

import (
	"context"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

type Employee struct {
	Id         int
	Name       string
	Department string
	Salary     int
}

func QueryEmployee() {
	db := ConnectDb()
	CreateTable(db, &Employee{})
	initEmploys(db)
	result := queryInBatch(db)
	if result != nil {
		fmt.Printf("查询到技术部记录数:%d \n", len(result))
	}
}

func initEmploys(db *gorm.DB) {
	//先删除所有的，然后重新插入
	gorm.G[Employee](db).Where("id>?", 0).Delete(context.Background())
	employees := make([]Employee, 0)
	for i := 1; i < 5; i++ {
		employees = append(employees, Employee{
			Id:         i,
			Name:       "employ_" + strconv.Itoa(i),
			Department: "技术部",
			Salary:     i * 10,
		})
	}
	gorm.G[Employee](db).CreateInBatches(context.Background(), &employees, len(employees))
}

func queryInBatch(db *gorm.DB) []Employee {
	if results, err := gorm.G[Employee](db).Where("department=?", "技术部").Find(context.Background()); err != nil {
		return nil
	} else {
		return results
	}
}
