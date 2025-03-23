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

	api := app.Group("/api/v1")
	users := api.Group("/users")

	users.Post("/", h.User.Create)
	users.Get("/", h.User.GetAll)
	users.Get("/:id", h.User.Get)
	users.Put("/:id", h.User.Update)
	users.Delete("/:id", h.User.Delete)
}
