package assign

import (
	"context"
	"encoding/json"
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
	agentRepo repository.Agent
	delayRepo repository.DelayReport
}

func New(cfg *config.Config, rabbitMQ *rabbitmq.Service, agentRepo repository.Agent, delayRepo repository.DelayReport) *Service {
	return &Service{
		cfg:       cfg,
		rabbitMQ:  rabbitMQ,
		agentRepo: agentRepo,
		delayRepo: delayRepo,
	}
}

func (s *Service) Handle(ctx context.Context, agentId uint) (api.Response, error) {
	agent, err := s.agentRepo.FindWithActiveDelayReportsById(ctx, agentId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return api.Response{
				Status:  http.StatusUnauthorized,
				Message: "Agent not found!",
			}, nil
		}

		return api.Response{}, err
	}

	if len(agent.DelayReports) > 0 {
		return api.Response{
			Status:  http.StatusForbidden,
			Message: "You have already assigned to an in-process delay report! Please review them first.",
		}, nil
	}

	var dr model.DelayReport

	msg, ok, err := s.rabbitMQ.Get()
	if err != nil {
		return api.Response{}, err
	}

	if !ok {
		return api.Response{
			Status:  http.StatusNotFound,
			Message: "No delivery report available",
		}, nil
	}

	err = json.Unmarshal(msg.Body, &dr)
	if err != nil {
		return api.Response{}, err
	}

	err = s.delayRepo.Assign(ctx, dr.ID, agentId)
	if err != nil {
		return api.Response{}, err
	}

	return api.Response{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Delivery report %v assigned to you", dr.ID),
	}, nil
}
