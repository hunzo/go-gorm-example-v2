package routes

import (
	"go-users/handlers"
	"go-users/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupRoutes(r *fiber.App) {
	db, err := gorm.Open(sqlite.Open("./db/db.sqlite"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	repo := services.NewUser(db)
	h := handlers.NewUserHandlers(repo)

	r.Get("/", h.GetUsers)
	r.Post("/add/user", h.AddUser)
	r.Get("/add/role/:rolename", h.AddRole)
}
