package common

import (
	"awesomeProject1/model"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/url"
)

var DB *gorm.DB

// InitDB 初始化数据库
func InitDB() *gorm.DB {
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc))
	//root:""@tcp(localhost:3306)/essential?charset=utf8&parseTime=True
	db, err := gorm.Open(mysql.Open(args))

	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Println("table created failed,err:", err)
	}

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
