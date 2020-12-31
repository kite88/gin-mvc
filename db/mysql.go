package db

import (
	"fmt"
	"gin-mvc/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

var Db *gorm.DB

func MysqlInit() {
	fmt.Println("启动数据库连接", time.Now().Format(conf.TimeFormatLayout))
	dbHost := conf.DbHost
	dbPort := conf.DbPort
	dbName := conf.DbName
	dbUser := conf.DbUser
	dbPassword := conf.DbPassword
	var err error
	Db, err = gorm.Open("mysql", dbUser+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln("连接数据库失败", err.Error())
	} else {
		fmt.Println("数据库连接成功", mysqlVersion(), time.Now().Format(conf.TimeFormatLayout))
	}
	Db.LogMode(true)
	Db.SingularTable(true)
}

func mysqlVersion() string {
	var rows string
	result := Db.Raw("select version() as version").Row()
	result.Scan(&rows)
	return "mysql " + rows
}
