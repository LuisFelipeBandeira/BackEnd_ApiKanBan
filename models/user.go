package models

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strings"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name" binding:"required,min=4,max=80"`
	Username string `json:"user" binding:"required,min=2,max=30"`
	Password string `json:"password" binding:"required,min=8"`
}

type UpdateUser struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"user"`
	Password string `json:"password"`
}

type LoginUser struct {
	Username string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user *User) EncriptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(user.Password))
	user.Password = hex.EncodeToString(hash.Sum(nil))
}

func (user *UpdateUser) EncriptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(user.Password))
	user.Password = hex.EncodeToString(hash.Sum(nil))
}

func (user *LoginUser) EncriptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(user.Password))
	user.Password = hex.EncodeToString(hash.Sum(nil))
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
