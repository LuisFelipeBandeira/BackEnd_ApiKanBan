package controllers

import (
	"net/http"
	"strconv"

	"github.com/LuisFelipeBandeira/BackEnd_ApiKanBan/models"
	"github.com/LuisFelipeBandeira/BackEnd_ApiKanBan/repositories"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	result, err := repositories.GetUsersRepository()
	if err != nil {
		messageError := err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"message": messageError})
		return
	}

	defer result.Close()

	var users []models.User

	for result.Next() {
		var user models.User

		if errScan := result.Scan(&user.ID, &user.Name, &user.Username, &user.Password); errScan != nil {
			messageError := errScan.Error()
			c.JSON(http.StatusInternalServerError, gin.H{"message": messageError})
			return
		}

		users = append(users, user)
	}

	c.JSON(200, users)
}

func NewUser(c *gin.Context) {
	var user *models.User

	if errBody := c.ShouldBindJSON(&user); errBody != nil {
		messageError := errBody.Error()
		c.JSON(http.StatusBadRequest, gin.H{"message": messageError})
		return
	}

	user.EncriptPassword()

	result, err := repositories.NewUserRepository(user)
	if err != nil {
		messageError := err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"message": messageError})
		return
	}

	lastId, errtoGetLastId := result.LastInsertId()
	if errtoGetLastId != nil {
		messageError := errtoGetLastId.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"message": messageError})
		return
	}

	c.JSON(200, gin.H{"message": "ID entered " + strconv.Itoa(int(lastId))})
}

func GetUserByID(c *gin.Context) {

	userId, errConvert := strconv.Atoi(c.Param("userid"))
	if errConvert != nil {
		messageError := errConvert.Error()
		c.JSON(http.StatusBadRequest, gin.H{"message": messageError})
		return
	}

}
