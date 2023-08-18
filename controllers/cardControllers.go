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

func GetCards(c *gin.Context) {
	cards, err := repositories.GetCardsRepository()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, cards)
}

func GetCardById(c *gin.Context) {

	cardId, errConvert := strconv.Atoi(c.Param("cardid"))
	if errConvert != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error converting given id"})
		return
	}

	card, err := repositories.GetCardByIdRepository(cardId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, card)
}

func DeleteCard(c *gin.Context) {
	cardId, errConvert := strconv.Atoi(c.Param("cardid"))
	if errConvert != nil {
		c.JSON(400, gin.H{"message": "error converting given id"})
		return
	}

	if err := repositories.DeleteCardRepository(cardId); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted successfully"})
}

func NewCard(c *gin.Context) {
	var card models.Card

	if errGetUser := c.ShouldBindJSON(&card); errGetUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "card submitted is invalid"})
		return
	}

	if errFormat := card.ValidAndFormat(); errFormat != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errFormat.Error()})
		return
	}

	token := strings.Split(c.GetHeader("Authorization"), " ")[1]

	userId, errGetUserId := services.GetUserIdByToken(token)
	if errGetUserId != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": errGetUserId.Error()})
		return
	}

	card.CreatedBy = uint(userId)

	card.CreatedAt = time.Now().Add(-time.Hour * 3)

	err := repositories.NewCardRepository(card)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "user created successfully"})
}

func FinishCard(c *gin.Context) {
	cardId, errConvert := strconv.Atoi(c.Param("cardid"))
	if errConvert != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error converting given id"})
		return
	}

	token := strings.Split(c.GetHeader("Authorization"), " ")[1]

	userId, errGetUserIdByToken := services.GetUserIdByToken(token)
	if errGetUserIdByToken != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": errGetUserIdByToken.Error()})
		return
	}

	user, errGetUser := repositories.GetUserByIDRepository(userId)
	if errGetUser != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": errGetUser.Error()})
		return
	}

	if err := repositories.FinishCardRepository(cardId, user, time.Now().Add(-time.Hour*3)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "the card was finished sucessfully"})
}

func UpdateCard(c *gin.Context) {
	cardId, errConvert := strconv.Atoi(c.Param("cardid"))
	if errConvert != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error converting given id"})
		return
	}

	userToken := strings.Split(c.GetHeader("Authorization"), " ")[1]

	userId, errGetUserIdByToken := services.GetUserIdByToken(userToken)
	if errGetUserIdByToken != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": errGetUserIdByToken.Error()})
		return
	}

	user, errGetUserById := repositories.GetUserByIDRepository(userId)
	if errGetUserById != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": errGetUserById.Error()})
		return
	}

	var CardFieldsToUpdate models.UpdateCard

	if errGetCardFields := c.ShouldBindJSON(&CardFieldsToUpdate); errGetCardFields != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errGetCardFields.Error()})
		return
	}

	if err := repositories.UpdateCardRepository(cardId, CardFieldsToUpdate, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "card updated successfully"})
}
