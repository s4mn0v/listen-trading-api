package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/s4mn0v/listen-trading-api/internal/api/handlers"
	"github.com/s4mn0v/listen-trading-api/internal/storage"
	"github.com/s4mn0v/listen-trading-api/logging/applogger"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	applogger.Info("Iniciando API de Copy-Trading en el puerto: " + port)

	storage.InitMongo()
	r := gin.Default()

	// 2 CORS ROUTES CONFIG
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-USER-KEY", "X-USER-SECRET", "X-USER-PASS"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Rutas
	traders := r.Group("/api/v2/traders")
	{
		traders.GET("/list", handlers.ListTraders)
		traders.GET("/detail/:id", handlers.TraderDetail)
	}

	r.Run(":" + port)
}
