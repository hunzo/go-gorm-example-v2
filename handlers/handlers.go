package handlers

import (
	"fmt"
	"go-users/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct {
	userRepo services.UserRepositoryDB
}

func NewUserHandlers(repo services.UserRepositoryDB) UserHandlers {
	return UserHandlers{userRepo: repo}
}

func (h UserHandlers) GetUsers(c *fiber.Ctx) error {
	users, err := h.userRepo.GetUses()
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}
	// getRoles

	for k, v := range users {
		fmt.Printf("%v, %v\n", k, v)
	}
	return c.JSON(fiber.Map{
		"users": users,
	})
}

func (h UserHandlers) AddUser(c *fiber.Ctx) error {
	type b struct {
		Username string `json:"username"`
	}
	req := b{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	user := services.Users{
		UserName: req.Username,
	}

	u, err := h.userRepo.AddUser(user)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"data":    "duplicate user",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "add user success",
		"data":    u,
	})
}

func (h UserHandlers) AddRole(c *fiber.Ctx) error {
	req := c.Params("rolename")
	err := h.userRepo.AddRole(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return c.JSON(fiber.Map{
		"success": true,
	})
}
