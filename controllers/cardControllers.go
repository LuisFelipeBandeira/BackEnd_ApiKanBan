package controllers

import (
	"net/http"
	"strconv"
	"strings"

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
		c.JSON(http.StatusBadRequest, gin.H{"message": "nao foi possivel converter o id informado"})
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

	sqlRow, err := repositories.GetUserByIDRepository(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var user models.User

	if errScan := sqlRow.Scan(&user); errScan != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": errScan.Error()})
		return
	}

	card.CreatedBy = user.Username
	card.Finished = 0

}
