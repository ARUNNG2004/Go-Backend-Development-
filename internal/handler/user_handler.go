package handler

import (
	"context"
	"database/sql"
	db "go-user-api/db/sqlc"
	"go-user-api/internal/logger"
	"go-user-api/internal/models"
	"go-user-api/internal/service"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UserHandler struct {
	DBQueries *db.Queries
	Validator *validator.Validate
}

// Create User
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.UserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	if err := h.Validator.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	parsedDob, _ := time.Parse("2006-01-02", req.Dob)

	id, err := h.DBQueries.CreateUser(context.Background(), db.CreateUserParams{
		Name: req.Name,
		Dob:  parsedDob,
	})
	if err != nil {
		logger.Log.Error("Failed to create user", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	logger.Log.Info("User created", zap.Int64("id", id))
	return c.Status(fiber.StatusCreated).JSON(models.UserResponse{
		ID:   id,
		Name: req.Name,
		Dob:  req.Dob,
	})
}

// Get User
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	user, err := h.DBQueries.GetUser(context.Background(), int64(id))
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	age := service.CalculateAge(user.Dob)

	return c.JSON(models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Format("2006-01-02"),
		Age:  age,
	})
}

// Update User
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var req models.UserRequest

	if err := c.BodyParser(&req); err != nil || h.Validator.Struct(req) != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	parsedDob, _ := time.Parse("2006-01-02", req.Dob)

	err := h.DBQueries.UpdateUser(context.Background(), db.UpdateUserParams{
		Name: req.Name,
		Dob:  parsedDob,
		ID:   int64(id),
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update"})
	}

	return c.JSON(models.UserResponse{
		ID:   int64(id),
		Name: req.Name,
		Dob:  req.Dob,
	})
}

// Delete User
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	err := h.DBQueries.DeleteUser(context.Background(), int64(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// List Users
func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.DBQueries.ListUsers(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	var response []models.UserResponse
	for _, u := range users {
		response = append(response, models.UserResponse{
			ID:   u.ID,
			Name: u.Name,
			Dob:  u.Dob.Format("2006-01-02"),
			Age:  service.CalculateAge(u.Dob),
		})
	}
	return c.JSON(response)
}
