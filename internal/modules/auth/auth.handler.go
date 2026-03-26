package auth

import (
	"github.com/doitung/DoiTung-service/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service AuthService
}

func NewAuthHandler(service AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Login(context *fiber.Ctx) error {

	var form LoginRequest

	if err := utils.ParseAndValidate(context, &form); err != nil {
		return utils.HandleError(context, err)
	}

	token, response, err := h.service.Login(form)
	if err != nil {
		return utils.HandleError(context, err)
	}

	context.Cookie(&fiber.Cookie{
		Name : "access_token",
		Value: token,
		HTTPOnly: true,
		Secure: false,
		SameSite: fiber.CookieSameSiteLaxMode,
		Path:     "/",
		MaxAge:   60 * 60 * 24,
	})

	return context.JSON(response)

}