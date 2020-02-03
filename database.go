package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type BlogDB struct {
	cnct *gorm.DB
}

func (it *BlogDB) Connect() {
	db, err := gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%v:%v/%v?parseTime=true",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_NAME"),
		),
	)
	if err != nil {
		BG.handleErrors(err.Error())
	}
	it.cnct = db
}

func (it *BlogDB) Close() {
	it.cnct.Close()
}
