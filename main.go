package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("Hello world\n")
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "pong ck!"})
	})
	err := r.Run()
	if err != nil {
		return
	}
}
