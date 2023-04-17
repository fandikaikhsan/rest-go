package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)


func getAlbums(c *gin.Context) {

	// setup database config
	ctx := context.Background()
	if err := godotenv.Load(); err != nil { panic(err) }
	dsn_url := os.Getenv("DATABASE_URL")

	dsn := dsn_url 
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	// setup redis

	// rdb_addr := os.Getenv("REDIS_URL")
	// rdb_pass := os.Getenv("REDIS_PASS")

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if err := rdb.Set(ctx, "key", "value", 0).Err() ; err != nil { panic(err) }

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil { panic(err) }

	fmt.Println("key: ", val)

	album := new(album)
	if err := db.NewSelect().Model(album).Where("id = ?", 1).Scan(ctx); err != nil { panic(err) }

	fmt.Printf("album: %#v", album)

	c.IndentedJSON(http.StatusOK, val)
	
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