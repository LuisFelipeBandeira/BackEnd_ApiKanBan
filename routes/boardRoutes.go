package routes

import (
	"github.com/LuisFelipeBandeira/BackEnd_ApiKanBan/controllers"
	"github.com/gin-gonic/gin"
)

func InitBoardRoutes(r *gin.RouterGroup) {
	r.GET("/boards", controllers.GetAllBoards)
	r.GET("/board/:boardid", controllers.GetBoardById)
}
