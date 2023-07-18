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

	//usecaseのコンストラクタを起動
	userUsecase := usecase.NewUserUsecase(userRepository)

	//controllerのコンストラクタを起動
	userController := controller.NewUserController(userUsecase)

	e := router.NewRouter(userController)

	e.Logger.Fatal(e.Start(":8080"))
}