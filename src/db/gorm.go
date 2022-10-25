package db

import (
	"fmt"
	"mini_e_commerce/src/server/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectGormDB() (*gorm.DB, error) {
	var (
		DB_HOST = "localhost"
		DB_PORT = "3306"
		DB_USER = "root"
		DB_PASS = ""
		DB_NAME = "mini_e_commerce"
	)

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	dbs, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = dbs.Ping()
	if err != nil {
		return nil, err
	}

	db.Debug().AutoMigrate(
		model.User{},
		model.Product{},
	)

	return db, nil
}
