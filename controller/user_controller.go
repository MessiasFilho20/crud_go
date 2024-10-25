package controller

import (
	"crud/models"
	"crud/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString("erro na requisição")
	}

	if err := service.CreateUser(user); err != nil {
		if err.Error() == "error" {
			return c.Status(400).SendString(err.Error())
		}
		return c.Status(500).SendString("Error ao criar user")
	}

	return c.JSON(user)
}

func GetUsers(c *fiber.Ctx) error {
	users, err := service.GetAllUsers()

	if err != nil {
		return c.Status(500).SendString("error ao obter usuarios")
	}
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(400).SendString("Id invalido")
	}

	user, err := service.GetUserById(uint(id))
	if err != nil {
		return c.Status(404).SendString("usuario não encontrado")
	}
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(400).SendString("Id invalido")
	}

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).SendString("error na requisição")
	}

	user.ID = uint(id)

	if err := service.UpdateUser(&user); err != nil {
		if err.Error() == "email ou cpf ja existe" {
			return c.Status(400).SendString(err.Error())
		}
		return c.Status(500).SendString("error atualizar usuarios")
	}

	return c.JSON(user)

}

func Deleteuser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(400).SendString("Id invalido")
	}

	if err := service.DeleteUser(uint(id)); err != nil {
		return c.Status(500).SendString("error ao deletar user")
	}
	return c.SendString("usuario deletado")
}
