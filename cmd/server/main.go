package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/immortal1405/vacay_planner/internal/config"
	"github.com/immortal1405/vacay_planner/internal/handlers"
	"github.com/immortal1405/vacay_planner/internal/services/llm"
	"github.com/immortal1405/vacay_planner/internal/services/planner"
	"github.com/immortal1405/vacay_planner/pkg/shivaay"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	shivaayClient := shivaay.NewClient(cfg.ShivaayAPIKey)

	langChainService := llm.NewLangChainService(shivaayClient)
	plannerService := planner.NewService(langChainService)
	handler := handlers.NewHandler(plannerService)

	r := gin.Default()
	r.POST("/api/v1/plan", handler.GeneratePlan)
	r.POST("/api/v1/tips", handler.GetTravelTips)
	r.POST("/api/v1/chat", handler.ChatWithPlanner)

	log.Printf("Starting server on :%s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
