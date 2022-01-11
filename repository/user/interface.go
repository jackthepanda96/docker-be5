package user

import "project/playground/be5/docker-play/entities"

type User interface {
	Get() ([]entities.User, error)
	Insert(newUser entities.User) (entities.User, error)
}
