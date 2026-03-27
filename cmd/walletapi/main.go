package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/sowens81/go-wallet-microservice/internal/application/wallet"
	"github.com/sowens81/go-wallet-microservice/internal/config"
	"github.com/sowens81/go-wallet-microservice/internal/delivery/api"
	"github.com/sowens81/go-wallet-microservice/internal/delivery/handler"
	"github.com/sowens81/go-wallet-microservice/internal/infrastructure/cosmos"
)

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("environment variable %s not set", key)
	}
	return value
}

func main() {
	// ---- Cosmos setup ----
	container, err := config.NewCosmosContainerWithAzureCredential(
		getEnv("COSMOS_ENDPOINT"),
		getEnv("COSMOS_DB"),
		getEnv("COSMOS_CONTAINER"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// ---- Dependency wiring ----
	repo := cosmos.NewWalletRepository(container)
	service := wallet.NewService(repo)

	walletHandler := handler.NewWalletHandler(service)
	walletRoutes := api.NewWalletRoutes(walletHandler)

	// ---- Gin router ----
	router := gin.Default()

	// ---- Register routes ----
	walletRoutes.Register(router)

	// ---- Start server ----
	log.Println("server running on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
