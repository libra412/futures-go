package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

// 全局数据库操作
var db *gorm.DB

// 初始化链接
func init() {
	DB, err := gorm.Open("mysql", "root:Btc123456!@tcp(150.158.126.196:3306)/btc?charset=utf8&parseTime=True&loc=Local")
	// defer db.Close()
	db = DB
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.LogMode(true)
	db.SingularTable(true)
	if err != nil {
		panic(err)
	}
	// 数据库头 拼接
	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	return "yz_" + defaultTableName
	// }

}

// K线
type Kline struct {
	Id             int64
	Open           string
	Close          string
	High           string
	Low            string
	Timestamp      string
	Volume         string
	CurrencyVolume string
	Type           int
	CreateTime     time.Time
}

// 添加K线
func (k *Kline) AddKline() bool {
	if db.NewRecord(k) {
		db.Create(k)
		return true
	}
	return false
}
