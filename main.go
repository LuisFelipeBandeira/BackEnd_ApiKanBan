package main

import (
	"Golang/BackEnd_ApiKanBan/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)
}
