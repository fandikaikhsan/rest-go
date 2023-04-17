package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getAlbums(secret, c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getRoot(c *gin.Context) {
	err := godotenv.Load()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	secret := os.Getenv("SECRET")
	c.IndentedJSON(http.StatusOK, secret)
}