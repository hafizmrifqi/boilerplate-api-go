package main

import (
	"boilerplate-api-go/internal/config"
	"boilerplate-api-go/internal/database"
	"boilerplate-api-go/internal/router"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg)
	defer db.Close()

	r := router.Setup()
	r.Run(":" + cfg.ServerPort)
}