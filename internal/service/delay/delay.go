package delay

import (
	"context"
	"fmt"
	"github.com/ashkan-jafarzadeh/delay/api"
	"github.com/ashkan-jafarzadeh/delay/config"
	"github.com/ashkan-jafarzadeh/delay/internal/model"
	"github.com/ashkan-jafarzadeh/delay/internal/repository"
	"github.com/ashkan-jafarzadeh/delay/pkg/rabbitmq"
	"gorm.io/gorm"
	"net/http"
)

type Service struct {
	cfg       *config.Config
	rabbitMQ  *rabbitmq.Service
	orderRepo repository.Order
	tripRepo  repository.Trip
	delayRepo repository.DelayReport
}

func New(cfg *config.Config, rabbitMQ *rabbitmq.Service, orderRepo repository.Order, tripRepo repository.Trip, delayRepo repository.DelayReport) *Service {
	return &Service{
		cfg:       cfg,
		rabbitMQ:  rabbitMQ,
		orderRepo: orderRepo,
		tripRepo:  tripRepo,
		delayRepo: delayRepo,
	}
}

func (s *Service) Handle(ctx context.Context, orderId uint) (api.Response, error) {
	order, err := s.orderRepo.FindById(ctx, orderId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return api.Response{
				Status:  http.StatusNotFound,
				Message: "Order not found!",
			}, nil
		}

		return api.Response{}, err
	}

	if !order.IsDelayed() {
		return api.Response{
			Status:  http.StatusUnprocessableEntity,
			Message: "Order is not delayed!",
		}, nil
	}

	if len(order.DelayReports) > 0 {
		return api.Response{
			Status:  http.StatusAlreadyReported,
			Message: "Order delayed has been already reported.",
		}, nil
	}

	dr := model.DelayReport{
		OrderId:  orderId,
		VendorId: order.VendorId,
	}

	err = s.delayRepo.Create(ctx, &dr)
	if err != nil {
		return api.Response{}, err
	}

	if order.Trip.IsActive() {
		newEstimate, err := findEstimate(s.cfg.Delivery.EstimateUrl)
		if err != nil {
			return api.Response{}, err
		}

		order.DeliveryTime = uint(newEstimate)
		err = s.orderRepo.Update(ctx, &order)
		if err != nil {
			return api.Response{}, err
		}

		return api.Response{
			Status:  http.StatusOK,
			Message: fmt.Sprintf("Sorry for the delay, New estimated time: %v minutes.", newEstimate),
		}, nil
	}

	go s.rabbitMQ.Publish(ctx, dr)

	return api.Response{
		Status:  http.StatusOK,
		Message: "Sorry for the delay, Our agents will review your report as soon as possible!",
	}, nil
}
