package handler

import (
	"errors"

	"github.com/canyouhearthemusic/gravitum/internal/domain/user"
	"github.com/canyouhearthemusic/gravitum/pkg/httpserver/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Router() *fiber.App {
	router := fiber.New()

	router.Get("/", h.getAll)
	router.Post("/", h.create)
	router.Get("/:id", h.get)
	router.Put("/:id", h.update)
	router.Delete("/:id", h.delete)

	return router
}

// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param request body user.CreateDTO true "User information"
// @Success 201 {object} user.Model
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /users [post]
func (h *UserHandler) create(c *fiber.Ctx) error {
	dto := new(user.CreateDTO)
	if err := c.BodyParser(dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.FailedPayload("invalid request body", nil))
	}

	if err := h.userService.CreateUser(c.Context(), dto); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.FailedPayload(err.Error(), nil))
	}

	return c.Status(fiber.StatusCreated).JSON(response.SuccessPayload("User has been successfully created!", nil))
}

// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Produce json
// @Success 200 {array} user.Model
// @Failure 500 {object} errorResponse
// @Router /users [get]
func (h *UserHandler) getAll(c *fiber.Ctx) error {
	users, err := h.userService.GetAllUsers(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.FailedPayload(err.Error(), nil))
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessPayload(nil, users))
}

// @Summary Get a user by ID
// @Description Get a user by their ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} user.Model
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /users/{id} [get]
func (h *UserHandler) get(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.FailedPayload("invalid user ID", nil))
	}

	user, err := h.userService.GetUser(c.Context(), id)
	if err != nil {
		if errors.Is(err, errors.New("user not found")) {
			return c.Status(fiber.StatusNotFound).JSON(response.FailedPayload("user not found", nil))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response.FailedPayload(err.Error(), nil))
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessPayload(nil, user))
}

// @Summary Update a user
// @Description Update a user's information
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param request body user.UpdateDTO true "User information to update"
// @Success 200 {object} user.Model
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /users/{id} [put]
func (h *UserHandler) update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.SuccessPayload("invalid user ID", nil))
	}

	req := new(user.UpdateDTO)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.FailedPayload("invalid request body", nil))
	}

	err = h.userService.UpdateUser(c.Context(), id, req)
	if err != nil {
		if errors.Is(err, errors.New("user not found")) {
			return c.Status(fiber.StatusNotFound).JSON(response.FailedPayload("user not found", nil))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response.FailedPayload(err.Error(), nil))
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessPayload("User has been successfully created!", nil))
}

// @Summary Delete a user
// @Description Delete a user by their ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 204 "No Content"
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /users/{id} [delete]
func (h *UserHandler) delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.FailedPayload("invalid user ID", nil))
	}

	err = h.userService.DeleteUser(c.Context(), id)
	if err != nil {
		if errors.Is(err, errors.New("user not found")) {
			return c.Status(fiber.StatusNotFound).JSON(response.FailedPayload("user not found", nil))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response.FailedPayload(err.Error(), nil))
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessPayload("User has been deleted!", nil))
}
