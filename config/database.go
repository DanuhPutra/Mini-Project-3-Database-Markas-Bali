package config

import (
	"DanuhPutra/MiniProject3/model"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDB struct {
	DB *gorm.DB
}

var Mysql MysqlDB

func OpenDB() {
	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=True",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"))
	MysqlConn, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	Mysql = MysqlDB{
		DB: MysqlConn,
	}

	err = autoMigrate(MysqlConn)
	if err != nil {
		log.Fatal(err)
	}
}

func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.DataBuku{},
	)
	if err != nil {
		return err
	}

	return nil
}