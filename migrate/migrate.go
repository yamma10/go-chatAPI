package main

import (
	"fmt"
	"go-chat-api/db"
	"go-chat-api/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.TalkRoom{}, &model.Message{})
	
	
}