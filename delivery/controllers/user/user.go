package user

import (
	"net/http"
	"project/playground/be5/docker-play/entities"
	userRepo "project/playground/be5/docker-play/repository/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type UserController struct {
	Repo userRepo.User
}

func New(user userRepo.User) *UserController {
	return &UserController{Repo: user}
}

func (uc UserController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := uc.Repo.Get()

		if err != nil {
			log.Info("Got error here")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Something wrong",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all data",
			"data":    users,
		})
	}
}

func (uc UserController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		requestFormat := CreateUserRequest{}

		if err := c.Bind(&requestFormat); err != nil {
			return c.JSON(http.StatusBadRequest, "There is something wrong with the input")
		}

		newUser := entities.User{
			Nama: requestFormat.Nama,
			HP:   requestFormat.HP,
		}

		res, err := uc.Repo.Insert(newUser)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "There is something wrong")
		}

		return c.JSON(http.StatusOK, res)
	}
}
