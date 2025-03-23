package handler

import (
	"github.com/canyouhearthemusic/gravitum/internal/service"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

type errorResponse struct {
	Error string `json:"error"`
}

type Handler struct {
	User *UserHandler
}

func New(services *service.Services) *Handler {
	return &Handler{
		User: NewUserHandler(services.User),
	}
}

func (h *Handler) Register(app *fiber.App) {
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	v1 := app.Group("/api/v1")
	v1.Mount("/users", h.User.Router())
}
