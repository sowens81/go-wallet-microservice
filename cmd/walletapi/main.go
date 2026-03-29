package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

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
	// load .env when present (ignore in production)
	// load local .env for development (ignored in prod)
	_ = godotenv.Load()

	endpoint := os.Getenv("COSMOS_ENDPOINT")
	db := os.Getenv("COSMOS_DB")
	containerName := os.Getenv("COSMOS_CONTAINER")
	tenantId := os.Getenv("AZURE_TENANT_ID")
	clientId := os.Getenv("AZURE_CLIENT_ID")
	clientSecret := os.Getenv("AZURE_CLIENT_SECRET")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	container, err := config.NewCosmosContainerWithAzureCredential(ctx, tenantId, clientId, clientSecret, endpoint, db, containerName)
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
