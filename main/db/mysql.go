package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// _ "github.com/go-sql-driver/mysql"
)


func ConnectMySQL(host, port, user, password, dbName string) (*gorm.DB, error){
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	

	if err != nil {
		return nil, err
	}

	return db, nil

}

