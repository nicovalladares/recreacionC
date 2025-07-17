package main

import (
	"fmt"
	"main/db"
	"main/transport"
	"net/http"
	"time"
)

func main(){

	db, err := db.ConnectMySQL("localhost", "3306", "root", "1234", "recreacionC") 
	if err != nil {
		fmt.Println("Error MySQL:", err)
		return
	}
	fmt.Println("Conectada a MySQL:", err)

	router := transport.NewRouter(db.DB)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	s.ListenAndServe()	

	// defer db.Close()



} 
