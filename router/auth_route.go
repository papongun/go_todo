package router

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	auth_c "github.com/papongun/go_todo/controller/auth"
	"github.com/papongun/go_todo/repository"
	auth_s "github.com/papongun/go_todo/service/auth"
)

func InitAuthRoute(router fiber.Router, db *gorm.DB) {
	userRepo := repository.GetUserRepository(db)
	authRegService := auth_s.GetAuthRegisterService(*userRepo)
	authRegController := auth_c.GetUserRegisterContoller(*authRegService)

	authRoute := router.Group("/auth")
	authRoute.Post("/login", authRegController.Register)
}
