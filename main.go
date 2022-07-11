package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-hello/version"
	"net/http"
	"time"
)

func main() {
	startTime := time.Now()
	appVersion := version.AppVersion()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("receive request at: ", time.Now())
		c.JSON(http.StatusOK, gin.H{
			"version":     appVersion,
			"startTime":   startTime,
			"currentTime": time.Now(),
		})
	})
	fmt.Println("server is ready to handle request...")
	r.Run(":8080")
}
