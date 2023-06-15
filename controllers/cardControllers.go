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
	sqlRows, err := repositories.GetCardsRepository()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	defer sqlRows.Close()

	var cards []models.Card

	for sqlRows.Next() {
		var card *models.Card

		if errScan := sqlRows.Scan(&card.ID, &card.Board, &card.Desc, &card.CreatedBy, &card.CreatedAt, &card.FinishedBy, &card.Finished, &card.FinishedAt); errScan != nil {
			c.JSON(500, gin.H{"message": errScan.Error()})
			return
		}

		cards = append(cards, *card)
	}

	c.JSON(200, cards)
}

func GetCardById(c *gin.Context) {

	cardId, errConvert := strconv.Atoi(c.Param("cardid"))
	if errConvert != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cardid enviado não pode ser convertido para int"})
		return
	}

	sqlRow, err := repositories.GetCardByIdRepository(cardId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var card *models.Card

	if errScan := sqlRow.Scan(&card.ID, &card.Board, &card.Desc, &card.CreatedBy, &card.CreatedAt, &card.FinishedBy, &card.Finished, &card.FinishedAt); errScan != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": errScan.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"message": "Card submitted is invalid"})
		return
	}

	if errFormat := card.ValidAndFormat(); errFormat != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errFormat.Error()})
		return
	}

	header := c.GetHeader("Authorization")
	if header == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(header, " ")[1]

	userId, errGetUserId := services.GetUserIdByToken(token)
	if errGetUserId != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": errGetUserId.Error()})
		return
	}

	sqlRow, errGetUserLogged := repositories.GetUserByIDRepository(userId)
	if errGetUserLogged != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": errGetUserLogged.Error()})
		return
	}

	var user models.User

	if errScan := sqlRow.Scan(&user.ID, &user.Name, &user.Username, &user.Password); errScan != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": errScan.Error()})
		return
	}

	card.CreatedBy = user.Username

	card.CreatedAt = time.Now()

	cardCadastrado, err := repositories.NewCardRepository(card)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, cardCadastrado)
}

func FinishCard(c *gin.Context) {
	cardId, errConvert := strconv.Atoi(c.Param("cardid"))
	if errConvert != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cardid enviado não pode ser convertido para int"})
		return
	}

	token := strings.Split(c.GetHeader("Authorization"), " ")[1]

	userId, errGetUserIdByToken := services.GetUserIdByToken(token)
	if errGetUserIdByToken != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": errGetUserIdByToken.Error()})
		return
	}

	var user models.User

	sqlRow, errGetUser := repositories.GetUserByIDRepository(userId)
	if errGetUser != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": errGetUser.Error()})
		return
	}

	if errScan := sqlRow.Scan(&user.ID, &user.Name, &user.Username, &user.Password); errScan != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": errScan.Error()})
		return
	}

	if err := repositories.FinishCardRepository(cardId, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.JSON(200, gin.H{"message": "the card was finished sucessfully"})
}

func UpdateCard(c *gin.Context) {
	cardId, errConvert := strconv.Atoi(c.Param("cardid"))
	if errConvert != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cardid enviado não pode ser convertido para int"})
		return
	}

	userToken := c.GetHeader("Authorization")
	if userToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "user is not authorized"})
		return
	}

	userId, errGetUserIdByToken := services.GetUserIdByToken(userToken)
	if errGetUserIdByToken != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error to get userId by authorization token"})
		return
	}

	var user models.User

	sqlRow, errGetUserById := repositories.GetUserByIDRepository(userId)
	if errGetUserById != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": errGetUserById.Error()})
		return
	}

	if errScanUser := sqlRow.Scan(&user.ID, &user.Name, &user.Username, &user.Password); errScanUser != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": errScanUser.Error()})
		return
	}

	if _, errGetCardById := repositories.GetCardByIdRepository(cardId); errGetCardById != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errGetCardById.Error()})
		return
	}

	var CardFieldsToUpdate models.UpdateCard

	if errGetCardFields := c.ShouldBindJSON(&CardFieldsToUpdate); errGetCardFields != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errGetCardFields.Error()})
	}

	if err := repositories.UpdateCardRepository(cardId, CardFieldsToUpdate, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "card updated successfully"})
}
