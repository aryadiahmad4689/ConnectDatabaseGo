package database

import (
	// "ConnectDatabaseGo"
	"ConnectDatabaseGo/ConnectDatabaseGo/utils"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Connect Database Func
func Connect(c string) *gorm.DB {
	// var DB Db
	if c == "pg" {
		pg()
	}
	db := mysql()
	return db

}

func pg() *gorm.DB {
	env := utils.LoadEnv()
	connection := "host=" + env.Host + " port=" + env.Port + " user=" + env.User + " dbname=" + env.Database + " password=" + env.Password + " sslmode=" + env.SSLMode
	db, err := gorm.Open("postgres", connection)
	if err != nil {
		log.Panic(err.Error())
	}
	return db
}

func mysql() *gorm.DB {
	env := utils.LoadEnv()
	connection := env.User + ":" + env.Password + "@" + (env.Host) + "/" + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", connection)
	if err != nil {
		log.Panic(err.Error())
	}
	// fmt.Print("anda berhasil connect")
	return db
}
