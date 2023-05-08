package models

import (
	"errors"
	"strings"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"user"`
	Password string `json:"password"`
}

func (user *User) ValidAndFormat() error {
	switch {
	case user.Name == "":
		return errors.New("o nome do usuario nao pode ser vazio")
	case user.Username == "":
		return errors.New("o username do usuario nao pode ser vazio")
	case user.Password == "":
		return errors.New("a senha do usuario nao pode ser vazia")
	default:
		user.Name = strings.TrimSpace(user.Name)
		user.Name = strings.ToLower(user.Name)
		user.Username = strings.TrimSpace(user.Username)
		user.Username = strings.ToLower(user.Username)
		return nil
	}
}
