package servers

import (
	_usersHttp "auth-example/modules/users/controllers"
	_usersRepository "auth-example/modules/users/repositories"
	_usersUsecases "auth-example/modules/users/usecases"

	_authHttp "auth-example/modules/auth/controllers"
	_authRepository "auth-example/modules/auth/repositories"
	_authUsecase "auth-example/modules/auth/usecases"

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

	//* Auth Group
	authGroup := v1.Group("/auth")
	authRepository := _authRepository.NewAuthRepository(s.Db)
	authUsecase := _authUsecase.NewAuthUsecase(authRepository, usersRepository)
	_authHttp.NewAuthController(authGroup, s.Cfg, authUsecase)

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
