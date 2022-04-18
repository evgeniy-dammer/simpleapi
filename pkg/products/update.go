package products

import (
	"github.com/evgeniy-dammer/simpleapi/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

type UpdateUserData struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func (h handler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UpdateUserData{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	user.Name = body.Name
	user.Phone = body.Phone
	user.Email = body.Email

	// save product
	h.DB.Save(&user)

	return c.Status(fiber.StatusOK).JSON(&user)
}
