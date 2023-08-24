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

	//usecaseのコンストラクタを起動
	userUsecase := usecase.NewUserUsecase(userRepository)
	roomUsecase := usecase.NewTalkRoomUsecase(roomRepository)

	//controllerのコンストラクタを起動
	userController := controller.NewUserController(userUsecase)
	roomController := controller.NewTalkRoomController(roomUsecase)

	e := router.NewRouter(userController, roomController)

	e.Logger.Fatal(e.Start(":8080"))
}