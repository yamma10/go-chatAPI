package main

import (
	"fmt"
	"go-chat-api/db"
	
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	//dbConn.AutoMigrate(&model.User{}, &model.Task{})
	
	
}