package task3

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Account struct {
	Id      int
	Balance int
}

type Transaction struct {
	Id            int
	FromAccountId int
	ToAccountId   int
	Amount        int
}

func DoTransaction() {
	db := ConnectDb()
	CreateTable(db, &Account{})
	CreateTable(db, &Transaction{})
	initRecord(db)
	accountAId := 1
	accountBId := 2
	tradeAmount := 300
	db.Transaction(func(tx *gorm.DB) error {
		var accountA Account
		var accountB Account
		var err error
		if accountA, err = gorm.G[Account](db).
			Where("id=?", accountAId).
			First(context.Background()); err != nil {
			fmt.Printf("查询A账户异常，交易将回滚,err:%s \n", err.Error())
			return err
		}
		if accountB, err = gorm.G[Account](db).
			Where("id=?", accountBId).
			First(context.Background()); err != nil {
			fmt.Printf("查询B账户异常，交易将回滚,err:%s \n", err.Error())
		}

		if tradeAmount > accountA.Balance {
			fmt.Printf("A账户余额%d不足，交易%d将回滚 \n", accountA.Balance, tradeAmount)
			return errors.New("账户余额不足")
		}

		if err = gorm.G[Transaction](db).
			Create(tx.Statement.Context, &Transaction{
				FromAccountId: accountAId,
				ToAccountId:   accountBId,
				Amount:        tradeAmount,
			}); err != nil {
			fmt.Printf("查询B账户异常，交易将回滚,err:%s \n", err.Error())
			return err
		}
		if _, err = gorm.G[Account](db).
			Where("id=?", accountA.Id).
			Update(tx.Statement.Context, "balance", accountA.Balance-tradeAmount); err != nil {
			fmt.Printf("更新A账户余额异常，交易将回滚,err:%s \n", err.Error())
			return err
		}

		if _, err = gorm.G[Account](db).
			Where("id=?", accountB.Id).
			Update(tx.Statement.Context, "balance", (accountB.Balance + tradeAmount)); err != nil {
			fmt.Printf("更新B账户余额异常，交易将回滚,err:%s \n", err.Error())
			return err
		}

		gorm.G[Account](db).Where("id=?", accountBId).Update(tx.Statement.Context, "balance", tradeAmount+accountB.Balance)
		fmt.Printf("交易成功 \n")

		return nil
	}, &sql.TxOptions{})

	if accountA, err := gorm.G[Account](db).Where("id=?", accountAId).First(context.Background()); err == nil {
		fmt.Printf("交易后A账户余额 account:%d \n", accountA.Balance)
	}

	if accountB, err := gorm.G[Account](db).Where("id=?", accountBId).First(context.Background()); err == nil {
		fmt.Printf("交易后B账户余额 account:%d \n", accountB.Balance)
	}

}

func initRecord(db *gorm.DB) {
	gorm.G[Account](db).Where("id in ?", []int{1, 2}).Delete(context.Background())
	gorm.G[Transaction](db).Where("id > ?", 0).Delete(context.Background())
	gorm.G[Account](db).Create(context.Background(), &Account{
		Id:      1,
		Balance: 200,
	})

	gorm.G[Account](db).Create(context.Background(), &Account{
		Id:      2,
		Balance: 100,
	})
}
