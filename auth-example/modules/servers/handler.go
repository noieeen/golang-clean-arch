package servers

import (
	_usersHttp "auth-example/modules/users/controllers"
	_usersRepository "auth-example/modules/users/repositories"
	_usersUsecases "auth-example/modules/users/usecases"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) MapHandlers() error {
	// Group a version
	v1 := s.App.Group("/v1")

	//* Users Group
	usersGroup := v1.Group("/users")
	usersRepository := _usersRepository.NewUsersRepository(s.Db)
	usersUsecases := _usersUsecases.NewUsersUsecase(usersRepository)
	_usersHttp.NewUsersController(usersGroup, usersUsecases)

	// End point not found response
	s.App.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     "error, end point not found",
			"result":      nil,
		})
	})

	return nil

}
