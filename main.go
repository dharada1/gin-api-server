package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yhirano55/gin-api-server/controllers"
	"reflect"
	"strconv"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello, gin!"})
	})

	router.GET("/user/:id", func(c *gin.Context) {
		v := c.Param("id")
		id, err := strconv.Atoi(v)
		if err != nil {
			c.JSON(400, err)
			return
		}
		if id <= 0 {
			c.JSON(400, gin.H{"message": "id should be greater than 0"})
			return
		}
		ctrl := controllers.NewUser()
		result := ctrl.Get(id)
		if result == nil || reflect.ValueOf(result).IsNil() {
			c.JSON(404, gin.H{"message": "not found"})
			return
		}

		c.JSON(200, result)
	})

	router.Run(":8080")
}
