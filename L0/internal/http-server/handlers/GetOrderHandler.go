package handlers

import (
	"demo_service/internal/models"
	"demo_service/internal/repository"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type GetOrderHandler interface {
	GetOrderById(orderUid string) (models.Order, error)
}

type Request struct {
	OrderUid string `json:"order_uid"`
}

type Response struct {
	Status string       `json:"status"`
	Error  string       `json:"error,omitempty"`
	Order  models.Order `json:"order,omitempty" validate:"required"`
}

func New(getOrderHandler GetOrderHandler, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.getOrder.new"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		orderUid := chi.URLParam(r, "order_uid")
		if orderUid == "" {
			log.Error("order_uid parameter is required")
			render.JSON(w, r, Response{
				Status: "Error",
				Error:  "order_uid parameter is required",
			})
			return
		}

		log.Info("order id received", slog.Any("request", orderUid))

		order, err := getOrderHandler.GetOrderById(orderUid)

		// TODO: переписать в свич кейс
		if errors.Is(err, repository.OrderNotFound) {
			log.Error("order not found: ", slog.Any("uid: ", orderUid))
			render.JSON(w, r, "failed to get order")
			return
		}
		if errors.Is(err, repository.DeliveryNotFound) {
			log.Error("delivery not found")
			render.JSON(w, r, "failed to get delivery")
			return
		}
		if errors.Is(err, repository.PaymentNotFound) {
			log.Error("payment not found")
			render.JSON(w, r, "failed to get payment")
			return
		}
		if errors.Is(err, repository.ItemsNotFound) {
			log.Error("items not found")
			render.JSON(w, r, "failed to get items")
			return
		}

		log.Info("order received", slog.Any("order: ", orderUid))

		render.JSON(w, r, Response{
			Status: "Ok",
			Order:  order,
		})
	}
}
