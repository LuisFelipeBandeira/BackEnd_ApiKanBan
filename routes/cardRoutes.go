package routes

import (
	"github.com/LuisFelipeBandeira/BackEnd_ApiKanBan/controllers"
	"github.com/gin-gonic/gin"
)

func InitCardRoutes(r *gin.RouterGroup) {
	r.GET("cards", controllers.GetCards)
	r.GET("cards/:cardid", controllers.GetCardById)
	r.DELETE("cards/:cardid", controllers.DeleteCard)
}
