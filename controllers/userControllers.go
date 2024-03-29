package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/LuisFelipeBandeira/BackEnd_ApiKanBan/models"
	"github.com/LuisFelipeBandeira/BackEnd_ApiKanBan/repositories"
	"github.com/LuisFelipeBandeira/BackEnd_ApiKanBan/services"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users, err := repositories.GetUsersRepository()
	if err != nil {
		messageError := err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"message": messageError})
		return
	}

	c.JSON(200, users)
}

func NewUser(c *gin.Context) {
	var user models.User

	token := strings.Split(c.GetHeader("Authorization"), " ")[1]

	userIdToken, errToGetIdByToken := services.GetUserIdByToken(token)
	if errToGetIdByToken != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": errToGetIdByToken})
		return
	}

	isAdm, errUserIsAdm := repositories.UserIsAdm(userIdToken)
	if errUserIsAdm != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": errUserIsAdm})
		return
	}

	if !isAdm {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "usuario nao possui permissao de adm"})
		return
	}

	if errBody := c.ShouldBindJSON(&user); errBody != nil {
		messageError := errBody.Error()
		c.JSON(http.StatusBadRequest, gin.H{"message": messageError})
		return
	}

	user.EncriptPassword()

	user.CreatedAt = time.Now().Add(-time.Hour * 3)

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

	var user models.User

	user, err := repositories.GetUserByIDRepository(userId)
	if err != nil {
		messageError := err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"message": messageError})
		return
	}

	user.EncriptPassword()

	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {

	token := strings.Split(c.GetHeader("Authorization"), " ")[1]

	userIdToken, errToGetIdByToken := services.GetUserIdByToken(token)
	if errToGetIdByToken != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": errToGetIdByToken})
		return
	}

	isAdm, errUserIsAdm := repositories.UserIsAdm(userIdToken)
	if errUserIsAdm != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": errUserIsAdm})
		return
	}

	if !isAdm {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "usuario nao possui permissao de adm"})
		return
	}

	userId, errConvert := strconv.Atoi(c.Param("userid"))
	if errConvert != nil {
		messageError := errConvert.Error()
		c.JSON(400, gin.H{"message": messageError})
		return
	}

	result, err := repositories.DeleteUserRepository(userId)
	if err != nil {
		messageError := err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"message": messageError})
		return
	}

	intAffectedRows, errGetQtdAffectedRows := result.RowsAffected()
	if errGetQtdAffectedRows != nil {
		messageError := errGetQtdAffectedRows.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"message": messageError})
		return
	}

	StringaffectedRows := strconv.Itoa(int(intAffectedRows))

	c.JSON(200, gin.H{"affectedRows": StringaffectedRows})
}

func UpdateUser(c *gin.Context) {

	token := strings.Split(c.GetHeader("Authorization"), " ")[1]

	userIdToken, errToGetIdByToken := services.GetUserIdByToken(token)
	if errToGetIdByToken != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": errToGetIdByToken})
		return
	}

	isAdm, errUserIsAdm := repositories.UserIsAdm(userIdToken)
	if errUserIsAdm != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": errUserIsAdm})
		return
	}

	if !isAdm {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "usuario nao possui permissao de adm"})
		return
	}

	userId, errToConvert := strconv.Atoi(c.Param("userid"))
	if errToConvert != nil {
		messageError := errToConvert.Error()
		c.JSON(http.StatusBadRequest, gin.H{"message": messageError})
		return
	}

	var userToUpdate models.UpdateUser

	if errToGetBody := c.ShouldBindJSON(&userToUpdate); errToGetBody != nil {
		messageError := errToGetBody.Error()
		c.JSON(http.StatusBadRequest, gin.H{"message": messageError})
		return
	}

	if err := repositories.UpdateUserRepository(userId, userToUpdate); err != nil {
		messageError := err.Error()
		c.JSON(http.StatusBadRequest, gin.H{"message": messageError})
		return
	}

	c.JSON(200, gin.H{"message": "usuario atualizado"})
}

func Login(c *gin.Context) {
	var userLogin models.LoginUser

	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userId, errLogin := repositories.LoginRepository(userLogin)
	if errLogin != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errLogin.Error()})
		return
	}

	token, errToCreateToken := services.CreateToken(userId)
	if errToCreateToken != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": errLogin.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User loged ", "token": token})
}
