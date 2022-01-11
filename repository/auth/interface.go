package auth

import "project/playground/be5/docker-play/entities"

type Auth interface {
	Login(name, hp string) (entities.User, error)
}
