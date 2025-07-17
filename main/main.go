package main

import (
	"fmt"
	"main/db"
)

func main(){

	db, err := db.ConnectMySQL("localhost", "3306", "root", "1234", "recreacionC") 

	if err != nil {
		fmt.Println("Error connecting to MySQL:", err)
		return
	}
	fmt.Println("conectada a  MySQL:", err)

	defer db.Close()


} 
