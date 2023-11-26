package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/ashkan-jafarzadeh/delay/api"
	"net/http"
)

func (s *Server) responseInternalError(c *fiber.Ctx, err error) error {
	log.Error(err)

	return s.response(c, api.Response{
		Status:  http.StatusInternalServerError,
		Message: "Something went wrong.",
	}, http.StatusInternalServerError)
}

func (s *Server) responseValidationError(c *fiber.Ctx, msg string) error {
	return s.response(c, api.Response{
		Status:  http.StatusUnprocessableEntity,
		Message: msg,
	}, http.StatusUnprocessableEntity)
}

func (s *Server) response(c *fiber.Ctx, res api.Response, status int) error {
	return c.Status(status).JSON(res)
}
