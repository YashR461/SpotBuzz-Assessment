package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	service "yash_rastogi/SpotBuzz-Assessment/service"
	db "yash_rastogi/SpotBuzz-Assessment/persistence"
)

func main() {

	fmt.Println("Hello. This is SpotBuzz Assessment !!")
	
	//Initializes a new instance of the Gin router, which is the core component of the Gin web framework in Go.
	router := gin.Default()

	db.Init()

	//Used to define a route for handling HTTP requests.
	router.POST("/players", service.AddPlayer)
	router.PUT("/players/:id", service.UpdatePlayerByID)
	router.DELETE("/players/:id", service.DeletePlayerByID)
	router.GET("/players", service.GetPlayers)
	router.GET("/players/rank/:val", service.GetPlayerByRank)
	router.GET("/players/random", service.GetRandomPlayer)

	//Used to start the HTTP server and make it listen on a specific port, in this case, port 8080
	router.Run("localhost:8080")
}
