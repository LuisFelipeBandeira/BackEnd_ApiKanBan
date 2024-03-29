package models

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID            int       `json:"id"`
	Name          string    `json:"name" binding:"required,min=4,max=100"`
	Username      string    `json:"user" binding:"required,min=2,max=80"`
	Email         string    `json:"email" binding:"required"`
	Password      string    `json:"password" binding:"required,min=6, max=100"`
	AdmPermission int8      `json:"adm_permission" binding:"required"`
	CreatedAt     time.Time `json:"created_at"`
}

type UpdateUser struct {
	Name          string `json:"name"`
	Username      string `json:"user"`
	Password      string `json:"password"`
	Email         string `json:"email"`
	AdmPermission int8   `json:"adm_permission"`
}

type LoginUser struct {
	Email    string `json:"email" binding:"required"`
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
		return errors.New("the user's name can not be empty")
	case user.Username == "":
		return errors.New("the username can not be empty")
	case user.Password == "":
		return errors.New("the user's password can not be empty")
	case user.Email == "":
		return errors.New("the user's email can not be empty")
	default:
		user.Name = strings.ToLower(strings.TrimSpace(user.Name))
		user.Username = (strings.TrimSpace(user.Username))
		user.Email = (strings.TrimSpace(user.Email))

		if errEmailFormat := checkmail.ValidateFormat(user.Email); errEmailFormat != nil {
			return errors.New("the email format is invalid")
		}

		return nil
	}
}
