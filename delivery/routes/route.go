package routes

import (
	"project/playground/be5/docker-play/delivery/controllers/auth"
	"project/playground/be5/docker-play/delivery/controllers/book"
	"project/playground/be5/docker-play/delivery/controllers/user"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, uc *user.UserController, ac *auth.AuthController, bc *book.BookController) {
	e.GET("/users", uc.Get())

	e.POST("/login", ac.Login())

	e.POST("/users", uc.Insert())

	e.POST("/books", bc.Create())
}
