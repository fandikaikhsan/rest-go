package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func main () {

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
	return
}

// func main() {
// 	router := gin.Default()

// 	router.GET("/albums", getAlbums)
// 	router.POST("/albums", postAlbums)
// 	router.GET("/", getRoot)

// 	router.Run("localhost:8080")
// }

