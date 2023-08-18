package main

import (
	"log"

	"github.com/LuisFelipeBandeira/BackEnd_ApiKanBan/middlewares"
	"github.com/LuisFelipeBandeira/BackEnd_ApiKanBan/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middlewares.Auth())

	routes.InitCardRoutes(&router.RouterGroup)
	routes.InitUserRoutes(&router.RouterGroup)
	routes.InitBoardRoutes(&router.RouterGroup)
	routes.InitColunRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err.Error())
	}
}
