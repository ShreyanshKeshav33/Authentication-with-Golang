package main

import (
	routes "go-jwt/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT") //getenv

	if port == "" {
		port = "8000"
	}
	router := gin.New() //creating a new instance of the Gin router

	//gin.Logger() is a middleware function provided by Gin that logs each HTTP request. It logs information such as the HTTP method, the path, the status code of the response, and the time taken to process the request
	router.Use(gin.Logger()) //adds the logger middleware to the Gin router. This means that for every incoming HTTP request, Gin will first execute the logger middleware, which will log information about the request, and then proceed to handle the request further.

	//we are adding two routes
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/ap1-1", func(c *gin.Context) { // whenwe use gin we dont have to use w and r rquest response handlers
		c.JSON(200, gin.H{"success": "Access granted"}) //gin.H refers to a map[string]interface{} type used in the Gin,
		//gin.H{"success":"Access granted"} creates a map with a string key "success" and a string value "Access granted". This map is then passed as the second argument to the c.JSON() function.
	})

	router.GET("/ap1-2", func(c *gin.Context) { // whenwe use gin we dont have to use w and r rquest response handlers
		c.JSON(200, gin.H{"success": "Access granted"}) //gin.H is for setting headers
	})

	router.Run(":" + port)
}
