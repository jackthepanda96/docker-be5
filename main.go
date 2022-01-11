package main

import (
	"project/playground/be5/docker-play/configs"
	"project/playground/be5/docker-play/delivery/controllers/auth"
	"project/playground/be5/docker-play/delivery/controllers/book"
	"project/playground/be5/docker-play/delivery/controllers/user"
	"project/playground/be5/docker-play/delivery/routes"
	"project/playground/be5/docker-play/utils"

	_authRepo "project/playground/be5/docker-play/repository/auth"
	_bookRepo "project/playground/be5/docker-play/repository/book"
	_userRepo "project/playground/be5/docker-play/repository/user"

	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	config := configs.GetConfig()

	db := utils.InitDB(config)
	userRepo := _userRepo.New(db)
	userController := user.New(userRepo)
	authRepo := _authRepo.New(db)
	authController := auth.New(authRepo)
	bookRepo := _bookRepo.New(db)
	bookController := book.New(bookRepo)

	e := echo.New()

	routes.RegisterPath(e, userController, authController, bookController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
