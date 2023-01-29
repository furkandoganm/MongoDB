package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"projects/APIWithMongoDB_2/dto"
	"projects/APIWithMongoDB_2/models"
	"projects/APIWithMongoDB_2/repository"
)

//go:generate mockgen -destination=../mocks/sevice/mockTodoService.go -package=services projects/APIWithMongoDB_2/services TodoService

type DefaultTodoService struct {
	Repo repository.TodoRepository
}

type TodoService interface {
	Insert(req models.Todo) (*dto.TodoDTO, error)
	GetAll() ([]models.Todo, error)
	Delete(id primitive.ObjectID) (*dto.TodoDTO, error)
}

func (t DefaultTodoService) Insert(req models.Todo) (*dto.TodoDTO, error) {
	var res dto.TodoDTO
	if len(req.Title) <= 2 {
		res.Status = false
		return &res, nil
	}

	result, err := t.Repo.Insert(req)
	if err != nil || !result {
		res.Status = result
		return &res, err
	}

	res = dto.TodoDTO{Status: result}
	return &res, nil
}

func (t DefaultTodoService) GetAll() ([]models.Todo, error) {
	result, err := t.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (t DefaultTodoService) Delete(id primitive.ObjectID) (*dto.TodoDTO, error) {
	var status dto.TodoDTO
	result, err := t.Repo.Delete(id)
	status.Status = result
	if err != nil || !result {
		return &status, err
	}
	return &status, nil
}

func NewTodoService(Repo repository.TodoRepository) DefaultTodoService {
	return DefaultTodoService{Repo: Repo}
}
