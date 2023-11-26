package rest

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ashkan-jafarzadeh/delay/config"
	"github.com/ashkan-jafarzadeh/delay/internal/service/assign"
	"github.com/ashkan-jafarzadeh/delay/internal/service/delay"
	"github.com/ashkan-jafarzadeh/delay/internal/service/report"
)

type Server struct {
	cfg           *config.Config
	app           *fiber.App
	delayService  *delay.Service
	assignService *assign.Service
	reportService *report.Service
}
type ErrorResponse struct {
	Message string            `json:"message"`
	Details map[string]string `json:"details"`
}

func New(cfg *config.Config, delayService *delay.Service, assignService *assign.Service, reportService *report.Service) *Server {
	return &Server{
		app:           fiber.New(),
		cfg:           cfg,
		delayService:  delayService,
		assignService: assignService,
		reportService: reportService,
	}
}

func (s *Server) Serve(ctx context.Context) error {
	s.SetupRoutes()

	return s.app.Listen(fmt.Sprintf(":%d", s.cfg.Http.Port))
}
