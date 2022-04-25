package main

import (
	"log"
	"micro/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "users"})
	})
	router.GET("/users", requests.GetUsers)
	router.POST("/users", requests.CreateUser)
	router.PUT("/users/:userid", requests.UpdateUser)
	router.GET("/users/:userid", requests.GetUserById)
	router.DELETE("/users/:userid", requests.DeleteUser)

	router.GET("/users/:userid/addresses", requests.GetAddress)
	router.POST("/users/:userid/addresses", requests.CreateAddress)
	router.PUT("/users/:userid/addresses/:addressid", requests.UpdateAddress)
	router.GET("/users/:userid/addresses/:addressid", requests.GetAddressById)
	router.DELETE("/users/:userid/addresses/:addressid", requests.DeleteAddress)

	log.Fatal(router.Run(":8081"))
}
