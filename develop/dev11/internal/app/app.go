package app

import (
	"dev11/config"
	"dev11/db"
	"dev11/internal/http/handlers/events"
	"dev11/internal/repository"
	"dev11/internal/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func Run() {
	// Создание логгера
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("failed to start logger")
	}

	// Загрузка конфигурации
	cfg, err := config.Load()
	if err != nil {
		logger.Fatal("failed to load config")
	}

	logger.Info("config loaded")

	// Создание подключения к БД
	db, err := db.NewPg(cfg.Db)
	if err != nil {
		logger.Fatal("failed to start db")
	}

	// Репозиторий событий
	repo := repository.NewPgEventRepository(db)
	// Сервис событий
	service := service.NewEventService(repo)
	// Обработчик запросов
	h := events.NewEventsHandler(service)
	r := h.SetupRoutes()

	// Запуск сервера
	go func() {
		if err := http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port), r); err != nil {
			logger.Fatal(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	
	logger.Info("Exit")
}
