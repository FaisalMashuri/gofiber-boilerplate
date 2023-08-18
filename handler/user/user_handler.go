package user

import (
	"github.com/gofiber/fiber/v2"
	"gofiber-boilerplate/internal/domain/user/model"
	"gofiber-boilerplate/internal/domain/user/model/user_dto"
	"gofiber-boilerplate/internal/domain/user/service"
	"net/http"
)

type UserHandler interface {
	//	Todo : Register function
	Create(ctx *fiber.Ctx) error
}

type UserHandlerImpl struct {
	service service.UserService
}

func NewUserHandler(service *service.UserService) UserHandler {
	return &UserHandlerImpl{
		service: *service,
	}
}

func (h *UserHandlerImpl) Create(ctx *fiber.Ctx) error {
	var userRequest user_dto.UserRequest
	var userModel model.User
	err := ctx.BodyParser(&userRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}
	userModel = userRequest.ToModeluser()
	err = h.service.Create(userModel)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON("success")
}
