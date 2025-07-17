package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)


func ConnectMySQL(host, port, user, password, dbName string) (*sql.DB, error){
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
	db, err := sql.Open("mysql", dsn)
	

	if err != nil {
		return nil, err
	}

	return db, nil

}

