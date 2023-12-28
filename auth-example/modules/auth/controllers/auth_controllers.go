package controllers

import (
	"auth-example/configs"
	"auth-example/modules/entities"
	"auth-example/pkg/middlewares"

	"github.com/gofiber/fiber/v2"
)

type authCon struct {
	Cfg     *configs.Configs
	AuthUse entities.AuthUsecase
}

func NewAuthController(r fiber.Router, cfg *configs.Configs, authUse entities.AuthUsecase) {
	controller := &authCon{
		Cfg:     cfg,
		AuthUse: authUse,
	}

	r.Post("/login", controller.Login)
	r.Get("auth-test", middlewares.JwtAuthentication(), controller.AuthTest)
}

func (h *authCon) Login(c *fiber.Ctx) error {
	req := new(entities.UsersCredentials)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}
	res, err := h.AuthUse.Login(h.Cfg, req)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     nil,
		"result":      res,
	})

}

func (h *authCon) AuthTest(c *fiber.Ctx) error {
	id := c.Locals("user_id")
	username := c.Locals("username")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"messages":    "",
		"result": fiber.Map{
			"id":       id,
			"username": username,
		},
	})
}
