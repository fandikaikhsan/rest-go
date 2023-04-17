package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func getAlbums(c *gin.Context) {

	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dsn_url := os.Getenv("DATABASE_URL")

	dsn := dsn_url 
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	album := new(album)
	if err := db.NewSelect().Model(album).Where("id = ?", 1).Scan(ctx); err != nil { panic(err) }

	fmt.Printf("album: %#v", album)

	c.IndentedJSON(http.StatusOK, album)
	
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