package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/medvedev-v/radiocontest/internal/handler"
	"github.com/medvedev-v/radiocontest/internal/repository"
	"github.com/medvedev-v/radiocontest/pkg/database"
)

func StartAndServe() {
	// DB Connect
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	// Repos init
	operatorRepo := repository.NewOperatorRepository(db)

	// Handlers init
	operatorHandler := handler.NewOperatorHandler(operatorRepo)

	// Gin router
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // для продакшена заменить на конкретные домены
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// API V1
	api := r.Group("/api/v1")
	{
		// Operator Endpoints
		api.POST("/operator", operatorHandler.CreateOperator)
		api.GET("/operator/:id", operatorHandler.GetOperator)
		api.PUT("/operator/:id", operatorHandler.UpdateOperator)
		api.DELETE("/operator/:id", operatorHandler.DeleteOperator)
	}

	// Server start
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
