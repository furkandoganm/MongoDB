package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http/httptest"
	services "projects/APIWithMongoDB_2/mocks/sevice"
	"projects/APIWithMongoDB_2/models"
	"testing"
)

var td TodoHandler

var mockService *services.MockTodoService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)

	mockService = services.NewMockTodoService(ctrl)

	td = TodoHandler{mockService}

	return func() {
		defer ctrl.Finish()
	}
}

func TestTodoHandler_GetAll(t *testing.T) {

	trd := setup(t)
	defer trd()

	router := fiber.New()
	router.Get("/api/todos", td.GetAll)

	var FakeDataForHandler = []models.Todo{
		{Id: primitive.NewObjectID(), Title: "Title 1", Content: "Content 1"},
		{Id: primitive.NewObjectID(), Title: "Title 2", Content: "Content 2"},
		{Id: primitive.NewObjectID(), Title: "Title 3", Content: "Content 3"},
	}

	mockService.EXPECT().GetAll().Return(FakeDataForHandler, nil)

	req := httptest.NewRequest("GET", "/api/todos", nil)

	resp, _ := router.Test(req, 1)

	assert.Equal(t, 200, resp.StatusCode)
}
