package book

import (
	"net/http"
	"project/playground/be5/docker-play/entities"
	"project/playground/be5/docker-play/repository/book"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	Repo book.Book
}

func New(repository book.Book) *BookController {
	return &BookController{Repo: repository}
}

func (bk BookController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		request := CreateBookRequest{}
		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, "There is something wrong with the input")
		}

		newBook := entities.Book{
			Title:     request.Title,
			Author_ID: uint(request.AuthorID),
		}

		res, err := bk.Repo.Insert(newBook)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "There is something wrong")
		}

		return c.JSON(http.StatusOK, res)
	}
}
