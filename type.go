package main

import "github.com/uptrace/bun"

type album struct {

	bun.BaseModel `bun:"album,alias:album"`

	ID    string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}