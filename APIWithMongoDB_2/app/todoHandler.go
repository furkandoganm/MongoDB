package app

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"projects/APIWithMongoDB_2/models"
	"projects/APIWithMongoDB_2/services"
)

type TodoHandler struct {
	Service services.TodoService
}

func (h TodoHandler) Insert(c *fiber.Ctx) error {
	var req models.Todo

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	result, err := h.Service.Insert(req)

	if err != nil || !result.Status {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(result)
}

func (h TodoHandler) GetAll(c *fiber.Ctx) error {
	result, err := h.Service.GetAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(result)
}

func (h TodoHandler) Delete(c *fiber.Ctx) error {
	query := c.Params("id")
	cnv, _ := primitive.ObjectIDFromHex(query)

	result, err := h.Service.Delete(cnv)
	if err != nil || !result.Status {
		log.Fatalln(err)
		//return c.Status(http.StatusBadRequest).JSON(fiber.Map{"State": false})
		return c.Status(http.StatusNotFound).JSON(err)
	}

	//return c.Status(http.StatusOK).JSON(fiber.Map{"State": true})
	return c.Status(http.StatusOK).JSON(result.Status)
}
