package configs

import (
	"fmt"
	"majoo/models/produk"
	"majoo/models/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type ConfigDB struct {
	DB_Username string
	DB_Password string
	DB_Host     string
	DB_Port     string
	DB_Database string
}

func InitConfigDB() ConfigDB {
	var configDB  = ConfigDB{
		DB_Username: "root",
		DB_Password: "",
		DB_Host: "localhost",
		DB_Port: "3306",
		DB_Database: "tes_majoo",
	}
	return configDB
}

func InitDB() {
	configDB := InitConfigDB()

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		configDB.DB_Username,
		configDB.DB_Password,
		configDB.DB_Host,
		configDB.DB_Port,
		configDB.DB_Database)

	var error error
	DB, error = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if error != nil {
		panic("Database failed connection : " + error.Error())
	}
	Migration()
}

func Migration()  {
	DB.AutoMigrate(&produk.Produk{},&user.User{})
}