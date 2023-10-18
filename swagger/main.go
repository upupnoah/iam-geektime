package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/upupnoah/iam-geektime/swagger/api"
	_ "github.com/upupnoah/iam-geektime/swagger/docs"
	"log"
	"net/http"
)

var users []*api.User

func main() {
	r := gin.Default()
	r.POST("/users", Create)
	r.GET("/users/:name", Get)
	log.Fatal(r.Run(":5555"))
}

// Create a user in memory.
func Create(c *gin.Context) {
	var user api.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "code": 10001})
		return
	}

	for _, u := range users {
		if u.Name == user.Name {
			c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("user %s already exist", user.Name), "code": 10001})
			return
		}
	}

	users = append(users, &user)
	c.JSON(http.StatusOK, user)
}

// Get return the detail information for a user.
func Get(c *gin.Context) {
	username := c.Param("name")
	for _, u := range users {
		if u.Name == username {
			c.JSON(http.StatusOK, u)
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("user %s not exist", username), "code": 10002})
}
