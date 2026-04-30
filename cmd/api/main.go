package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/s4mn0v/listen-trading-api/internal/api/handlers"
	"github.com/s4mn0v/listen-trading-api/logging/applogger"
)

func main() {
	// Cargar archivo .env
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	applogger.Info("Iniciando API de Copy-Trading en el puerto: " + port)

	r := gin.Default()

	// Rutas de Traders
	traders := r.Group("/api/v2/traders")
	{
		traders.GET("/list", handlers.ListTraders)
	}

	r.Run(":" + port)
}
