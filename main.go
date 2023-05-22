package main

import (
	"log"

	"github.com/LuisFelipeBandeira/BackEnd_ApiKanBan/middlewares"
	"github.com/LuisFelipeBandeira/BackEnd_ApiKanBan/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middlewares.Auth())

	routes.InitCardRoutes(&router.RouterGroup)
	routes.InitUserRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err.Error())
	}
}
