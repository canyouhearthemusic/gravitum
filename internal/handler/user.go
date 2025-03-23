package handler

import (
	"errors"

	"github.com/canyouhearthemusic/gravitum/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type userCreateRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type userUpdateRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserHandler struct {
	userService service.UserServiceInterface
}

func NewUserHandler(userService service.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param request body userCreateRequest true "User information"
// @Success 201 {object} entity.User
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /users [post]
func (h *UserHandler) Create(c *fiber.Ctx) error {
	var req userCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Error: "invalid request body"})
	}

	user, err := h.userService.CreateUser(c.Context(), req.Username, req.Email, req.FirstName, req.LastName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Error: err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Produce json
// @Success 200 {array} entity.User
// @Failure 500 {object} errorResponse
// @Router /users [get]
func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.userService.GetAllUsers(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Error: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

// @Summary Get a user by ID
// @Description Get a user by their ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} entity.User
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /users/{id} [get]
func (h *UserHandler) Get(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Error: "invalid user ID"})
	}

	user, err := h.userService.GetUser(c.Context(), id)
	if err != nil {
		if errors.Is(err, errors.New("user not found")) {
			return c.Status(fiber.StatusNotFound).JSON(errorResponse{Error: "user not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Error: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

// @Summary Update a user
// @Description Update a user's information
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param request body userUpdateRequest true "User information to update"
// @Success 200 {object} entity.User
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /users/{id} [put]
func (h *UserHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Error: "invalid user ID"})
	}

	var req userUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Error: "invalid request body"})
	}

	user, err := h.userService.UpdateUser(c.Context(), id, req.Username, req.Email, req.FirstName, req.LastName)
	if err != nil {
		if errors.Is(err, errors.New("user not found")) {
			return c.Status(fiber.StatusNotFound).JSON(errorResponse{Error: "user not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Error: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(user)
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
func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{Error: "invalid user ID"})
	}

	err = h.userService.DeleteUser(c.Context(), id)
	if err != nil {
		if errors.Is(err, errors.New("user not found")) {
			return c.Status(fiber.StatusNotFound).JSON(errorResponse{Error: "user not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{Error: err.Error()})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}
