package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/", getRoot)

	router.Run("localhost:8080")
}

