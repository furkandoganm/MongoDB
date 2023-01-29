package services

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"projects/APIWithMongoDB_2/mocks/repository"
	"projects/APIWithMongoDB_2/models"
	"testing"
)

var mockRepo *repository.MockTodoRepository
var service TodoService

var FakeData = []models.Todo{
	{Id: primitive.NewObjectID(), Title: "Title 1", Content: "Content 1"},
	{Id: primitive.NewObjectID(), Title: "Title 2", Content: "Content 2"},
	{Id: primitive.NewObjectID(), Title: "Title 3", Content: "Content 3"},
}

func setup(t *testing.T) func() {
	ct := gomock.NewController(t)
	defer ct.Finish()
	mockRepo = repository.NewMockTodoRepository(ct)
	service = NewTodoService(mockRepo)
	return func() {
		service = nil
		defer ct.Finish()
	}
}

func TestDefaultTodoService_GetAll(t *testing.T) {
	td := setup(t)
	defer td()

	mockRepo.EXPECT().GetAll().Return(FakeData, nil)
	result, err := service.GetAll()

	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, result)
}
