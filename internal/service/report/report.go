package report

import (
	"context"
	"github.com/ashkan-jafarzadeh/delay/api"
	"github.com/ashkan-jafarzadeh/delay/config"
	"github.com/ashkan-jafarzadeh/delay/internal/aggregate"
	"github.com/ashkan-jafarzadeh/delay/internal/repository"
	"net/http"
	"time"
)

type Service struct {
	cfg        *config.Config
	vendorRepo repository.Vendor
}

func New(cfg *config.Config, vendorRepo repository.Vendor) *Service {
	return &Service{
		cfg:        cfg,
		vendorRepo: vendorRepo,
	}
}

type Output struct {
	Data []aggregate.VendorReport
}

func (s *Service) Handle(ctx context.Context) (api.Response, error) {
	result, err := s.vendorRepo.Report(ctx, time.Now().AddDate(0, 0, -7))
	if err != nil {
		return api.Response{}, err
	}
	return api.Response{
		Data:   result,
		Status: http.StatusOK,
	}, nil
}
