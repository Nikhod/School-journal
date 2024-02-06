package handlers

import (
	"go.uber.org/zap"
	"second/internal/services"
)

type Handler struct {
	Service *services.Services
	Logger  *zap.Logger
}

func NewHandler(service *services.Services, logger *zap.Logger) *Handler {
	return &Handler{Service: service, Logger: logger}
}
