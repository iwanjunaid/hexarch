package v1

import (
	"github.com/iwanjunaid/hexarch/internal/core/port/inbound/registry"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	logger          *log.Entry
	serviceRegistry registry.ServiceRegistry
}

func New(logger *log.Entry, serviceRegistry registry.ServiceRegistry) *Handler {
	return &Handler{
		logger:          logger,
		serviceRegistry: serviceRegistry,
	}
}

func (h *Handler) GetLogger() *log.Entry {
	return h.logger
}

func (h *Handler) GetServiceRegistry() registry.ServiceRegistry {
	return h.serviceRegistry
}
