package main

import (
	"fmt"
	"log"
	"net/http"

	"gosolve/backend/internal/api"
	"gosolve/backend/internal/controller"
	"gosolve/backend/internal/repository"
	"gosolve/backend/internal/services"
	"gosolve/backend/internal/utils"

	"go.uber.org/zap"
)

func main() {

	//logger, _ := utils.NewLogger()
	config := utils.LoadConfig()
	logger, err := utils.NewCustomLogger()

	if err != nil {
		log.Fatalf("Failed to initialize logger: %s\n", err)
	}
	defer logger.Sync()

	repo, err := repository.NewDataRepository(logger)
	if err != nil {
		logger.Fatal("Failed to initialize repository", zap.Error(err))
	}

	if repo == nil {
		logger.Fatal("Repository instance is nil, exiting...")
	}

	service := services.NewSearchService(repo, logger)
	controller, _ := controller.New(controller.Params{SearchService: service})

	handler := api.NewHandler(controller)
	router := api.SetupRouter(handler)

	logger.Info("Server is starting...",
		zap.String("port", config.ServerPort),
		zap.String("log_level", config.LogLevel),
	)
	addr := fmt.Sprintf(":%s", config.ServerPort)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}

}
