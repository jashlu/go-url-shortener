package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	//create a new Gin router which will handle all incoming HTTP requests
	r := gin.Default()
	//sets up an HTTP GET route at the path "/"
	//So when someone visits http://localhost:9808, this handler will execute
	r.GET("/hello", func(c *gin.Context) {
		//this will respond with a JSON object containing the following message and 200 HTTP code
		c.JSON(200, gin.H{
			"message": "Hey Go URL Shortener !",
		})
	})

	//this starts the web server and listens on port 9808
	err := r.Run(":9808")

	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
