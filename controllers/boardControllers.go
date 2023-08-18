package controllers

import (
	"net/http"

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
