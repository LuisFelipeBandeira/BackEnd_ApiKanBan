package routes

import (
	"github.com/LuisFelipeBandeira/BackEnd_ApiKanBan/controllers"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.RouterGroup) {
	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.NewUser)
}
