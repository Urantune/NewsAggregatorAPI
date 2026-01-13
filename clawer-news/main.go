package main

import (
	"clawer-news/headlers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/test", headlers.GetNews)

	r.Run(":8081")
}
