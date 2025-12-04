package task3

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Student struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Age       int
	Grade     string
}

func Curd() {
	db := ConnectDb()
	CreateTable(db, &Student{})
	createRecord(db)
	queryAgeGrante18(db)
	updateRecord(db)
	delteRecord(db)
}

func ConnectDb() *gorm.DB {
	// 1. 配置 GORM 日志：打印最终执行的 SQL 到控制台
	newLogger := logger.New(
		// 将日志输出到控制台（os.Stdout）
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值（超过 1 秒会标红）
			LogLevel:                  logger.Info, // 日志级别：Info 会打印所有 SQL
			IgnoreRecordNotFoundError: true,        // 忽略记录不存在的错误
			Colorful:                  true,        // 开启彩色输出（控制台更易读）
		},
	)
	dsn := "root:mysql@tcp(127.0.0.1:33066)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	return db
}

func updateRecord(db *gorm.DB) {
	affected, _ := gorm.G[Student](db).Where("name=?", "张三").Update(context.Background(), "grade", "四年级")
	fmt.Printf("updateRecord,affected rows:%d \n", affected)

	_, _ = gorm.G[Student](db).Where("name=? and id<?", "张三", 5).Updates(context.Background(), Student{
		Age:   10,
		Grade: "三年级",
	})

}

func delteRecord(db *gorm.DB) {
	//物理删除
	affected, _ := gorm.G[Student](db.Unscoped()).Where("id<?", 5).Delete(context.Background())
	fmt.Printf("delete records,affected rows:%d \n", affected)
}

func CreateTable(db *gorm.DB, dst interface{}) {

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

func createRecord(db *gorm.DB) {
	var student = Student{
		Name:  "张三",
		Age:   20,
		Grade: "三年级",
	}
	fmt.Println("开始创建记录")
	tx := db.Create(&student)
	fmt.Println("创建记录结束")
	fmt.Printf("insert rows:%d,id:%d \n", tx.RowsAffected, student.ID)
}

func queryAgeGrante18(db *gorm.DB) {
	result, err := gorm.G[Student](db).Where("age>?", 18).Find(context.Background())
	if err != nil {
		panic(err)
	}
	b, _ := json.Marshal(result)
	fmt.Printf("找到记录数:%d,val:%s \n", len(result), string(b))
}
