package auth

import "project/playground/be5/docker-play/entities"

type LoginRequestFormat struct {
	Name string `json:"name" form:"name"`
	HP   string `json:"hp" form:"hp"`
}

type LoginResponseFormat struct {
	Message string        `json:"message"`
	Data    entities.User `json:"data"`
	Token   string        `json:"token"`
}
