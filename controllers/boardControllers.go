package controllers

import (
	"net/http"
	"strconv"

	"github.com/LuisFelipeBandeira/BackEnd_ApiKanBan/repositories"
	"github.com/gin-gonic/gin"
)

func GetAllBoards(c *gin.Context) {
	boards, errToGetBoards := repositories.GetAllBoards()
	if errToGetBoards != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": errToGetBoards.Error()})
		return
	}

	c.JSON(200, boards)
}

func GetBoardById(c *gin.Context) {
	boardId, errToConvertBoardid := strconv.Atoi(c.Param("boardid"))
	if errToConvertBoardid != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error to convert sent boardid to number"})
		return
	}

	board, codeError, errGetBoardById := repositories.GetBoardById(uint(boardId))
	if errGetBoardById != nil {
		c.JSON(codeError, gin.H{"message": errGetBoardById})
		return
	}

	c.JSON(200, board)
}
