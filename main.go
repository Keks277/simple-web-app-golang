package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	router.LoadHTMLGlob("html/*.html")
	router.GET("/ping", handlerPing)
	router.GET("/", handlerIndex)
	router.Run("127.0.0.1:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func handlerPing(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func handlerIndex(context *gin.Context) {
	context.HTML(200, "index.html", nil)
}
