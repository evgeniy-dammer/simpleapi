package products

import (
	"github.com/evgeniy-dammer/simpleapi/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

type AddUserData struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func (h handler) AddUser(c *fiber.Ctx) error {
	body := AddUserData{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var user models.User

	user.Name = body.Name
	user.Phone = body.Phone
	user.Email = body.Email

	if result := h.DB.Create(&user); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&user)
}
