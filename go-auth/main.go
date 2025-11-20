
package main

import (
  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()

  router.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "Welcome to the Go Authentication and Authorization tutorial!",
    })
  })

  
	router.POST("/register", func(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	c.JSON(200, gin.H{"message": "User registered successfully"})
	})


  router.Run()
}
