package main

import (
	"go.uber.org/zap"
	"log"
	"net/http"
	"second/internal/configs"
	"second/internal/handlers"
	"second/internal/repositories"
	"second/internal/services"
	"second/loggers"
	"second/pkg/database"
)

func main() {
	logger, err := loggers.InitLogger()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func(logger2 *zap.Logger) {
		err = logger2.Sync()
		if err != nil {
			log.Println(err)
			return
		}
	}(logger)

	execute(logger)
}

func execute(logger *zap.Logger) {

	config, err := configs.InitConfigs()
	if err != nil {
		logger.Error("Error in configs")
		return
	}

	db, err := database.InitDatabase(config)
	if err != nil {
		logger.Error("Error in database")
		return
	}

	repository := repositories.NewRepository(db, logger)
	service := services.NewService(repository, logger)
	handler := handlers.NewHandler(service, logger)
	mux := InitMux(handler)

	srv := http.Server{
		Addr:    config.Server.Host + config.Server.Port,
		Handler: mux,
	}
	err = srv.ListenAndServe()
	if err != nil {
		logger.Error("Issue in ListeningAndServe!", zap.Error(err))
		return
	}

}
