package book

import "project/playground/be5/docker-play/entities"

type Book interface {
	Insert(newUser entities.Book) (entities.Book, error)
}
