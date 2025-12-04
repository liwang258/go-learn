package task3

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

type Book struct {
	Id     int     `gorm:"column:id;type:bigint"`
	Title  string  `gorm:"type:varchar(100)"`
	Author string  `gorm:"column:author;type:varchar(30)"`
	Price  float32 `gorm:"column:price;type:decimal(5,2)"`
}

func TypeSafe() {
	db := ConnectDb()
	CreateTable(db, &Book{})
	initV := make([]Book, 0)
	for i := 0; i < 4; i++ {
		initV = append(initV, Book{
			Title:  "title_" + strconv.Itoa(i),
			Author: "author_" + strconv.Itoa(i),
			Price:  float32((i + 10*3)),
		})
	}
	gorm.G[Book](db).CreateInBatches(context.Background(), &initV, len(initV))

	if result, err := gorm.G[Book](db).Where("price>?", float32(30)).Find(context.Background()); err == nil {
		b, _ := json.Marshal(result)
		fmt.Printf("查找到记录:\n%s \n", string(b))
	}

}
