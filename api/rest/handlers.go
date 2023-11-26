package rest

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

func (s *Server) Delay(c *fiber.Ctx) error {
	orderId, err := c.ParamsInt("orderId")
	if err != nil {
		return s.responseValidationError(c, "Order ID is not valid")
	}

	res, err := s.delayService.Handle(c.Context(), uint(orderId))
	if err != nil {
		return s.responseInternalError(c, err)
	}

	return s.response(c, res, res.Status)
}

func (s *Server) Assign(c *fiber.Ctx) error {
	agentId, err := strconv.Atoi(c.Get("agent-id"))
	if err != nil {
		return err
	}
	if agentId == 0 || err != nil {
		return s.responseValidationError(c, "header key `agent-id` is not valid")
	}
	res, err := s.assignService.Handle(c.Context(), uint(agentId))
	if err != nil {
		return s.responseInternalError(c, err)
	}

	return s.response(c, res, res.Status)
}

func (s *Server) VendorReport(c *fiber.Ctx) error {
	res, err := s.reportService.Handle(c.Context())
	if err != nil {
		return s.responseInternalError(c, err)
	}

	return s.response(c, res, http.StatusOK)
}
