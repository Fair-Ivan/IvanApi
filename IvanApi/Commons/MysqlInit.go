package Commons

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

// 新建连接
func NewConn() *gorm.DB {
	connArgs := GetConfigJson().ConnectionString
	db, err := gorm.Open("mysql", connArgs)
	if err != nil {
		fmt.Println("%s", err.Error())
		panic("连接数据库失败")
	}
	return db
}

// 初始化gorm
func GormInit() {
	db = NewConn()
	db.DB().SetMaxIdleConns(30)   // 最大空闲连接数
	db.DB().SetMaxOpenConns(2000) // 最大连接数
}

// 获取连接
func Getdb() *gorm.DB {
	if err := db.DB().Ping(); err != nil {
		db.Close()
		db = NewConn()
	}
	return db
}
