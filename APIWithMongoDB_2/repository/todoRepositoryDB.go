package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"projects/APIWithMongoDB_2/models"
	"time"
)

//go:generate mockgen -destination=../mocks/repository/mockTodoRepository.go -package=repository projects/APIWithMongoDB_2/repository TodoRepository

type TodoRepositoryDB struct {
	TodoCollection *mongo.Collection
}

type TodoRepository interface {
	Insert(req models.Todo) (bool, error)
	GetAll() ([]models.Todo, error)
	Delete(id primitive.ObjectID) (bool, error)
}

func (t TodoRepositoryDB) Insert(req models.Todo) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req.Id = primitive.NewObjectID()
	result, err := t.TodoCollection.InsertOne(ctx, req)
	if result.InsertedID == nil || err != nil {
		errors.New("failed add")
		return false, err
	}
	return true, nil
}

func (t TodoRepositoryDB) GetAll() ([]models.Todo, error) {
	var req models.Todo
	var reqs []models.Todo

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.TodoCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	for result.Next(ctx) {
		if err := result.Decode(&req); err != nil {
			log.Fatalln(err)
			return nil, err
		}
		reqs = append(reqs, req)
	}
	return reqs, nil
}

func (t TodoRepositoryDB) Delete(id primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.TodoCollection.DeleteOne(ctx, bson.M{"id": id})

	if err != nil || result.DeletedCount <= 0 {
		return false, err
	}

	return true, nil
}

func NewTodoRepositoryDB(dbClient *mongo.Collection) TodoRepositoryDB {
	return TodoRepositoryDB{
		TodoCollection: dbClient,
	}
}
