package kafka

import (
	"demo_service/internal/models"
	"demo_service/internal/repository"
	"encoding/json"
	"fmt"
	"log/slog"
)

type ConsumerHandler struct {
	repository *repository.Repository
	log        *slog.Logger
}

func (h *ConsumerHandler) HandleMessage(msg []byte) error {
	var order models.Order
	if err := json.Unmarshal(msg, &order); err != nil {
		return fmt.Errorf("msg unmarshaling error: %w", err)
	}

	h.log.Info("received order: ", order.OrderUId)

	if err := h.repository.SaveOrder(order); err != nil {
		return fmt.Errorf("saving message error: %w", err)
	}

	h.log.Info("msg successfully sent")

	return nil
}

func NewHandler(repo *repository.Repository, logger *slog.Logger) *ConsumerHandler {
	return &ConsumerHandler{repository: repo, log: logger}
}
