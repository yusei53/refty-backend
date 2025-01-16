package main

import (
	_ "refty-backend/docs" // Swagger docs の読み込み
	"refty-backend/internal/app"
)

// @title Refty API
// @version 1.0
// @description This is the Refty API server.

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	a, err := app.New()
	if err != nil {
		panic(err)
	}

	if err := a.Run(); err != nil {
		panic(err)
	}
}
