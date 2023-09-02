package main

import (
	"go-chat-api/controller"
	"go-chat-api/db"
	"go-chat-api/repository"
	"go-chat-api/router"
	"go-chat-api/usecase"
)

func main() {
	db := db.NewDB()
	//repositoryのコンストラクタを起動
	userRepository := repository.NewUserRepository(db)
	roomRepository := repository.NewTalkRoomRepository(db)
	messageRepository := repository.NewMessageRepository(db)

	//usecaseのコンストラクタを起動
	userUsecase := usecase.NewUserUsecase(userRepository)
	roomUsecase := usecase.NewTalkRoomUsecase(roomRepository)
	messageUsecase := usecase.NewMessageUsecase(messageRepository)

	//controllerのコンストラクタを起動
	userController := controller.NewUserController(userUsecase)
	roomController := controller.NewTalkRoomController(roomUsecase)
	messageController := controller.NewMessageController(messageUsecase)

	e := router.NewRouter(userController, roomController, messageController)

	e.Logger.Fatal(e.Start(":8080"))
}