package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var DB *gorm.DB

func Database(conn string) {
	//fmt.Println(conn)
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		panic(err)
	}
	fmt.Println("MySQL数据库连接成功")
	db.LogMode(true)
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	db.SingularTable(true) // 表格名字不加s，默认gorm会写成users
	sqlDB := db.DB()
	sqlDB.SetMaxIdleConns(20)  //设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) //打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	DB = db
	migration()
}
