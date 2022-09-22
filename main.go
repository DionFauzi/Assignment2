package main

import (
	config "ASSIGNMENT2/configs"
	database "ASSIGNMENT2/helpers/database"
	routerConfig "ASSIGNMENT2/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.InitConfig()
	database.StartDB(config.Database)

	var PORT = config.WebServer.Port

	routers := gin.Default()
	routerConfig.RouteOrders(routers)

	routers.Run(PORT)
}
