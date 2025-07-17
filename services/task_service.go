package services

import (
	"github.com/andrey-lawyer/go-gin-todo-app/models"
	"github.com/andrey-lawyer/go-gin-todo-app/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskService struct {
	TaskRepo *repositories.TaskRepository
}

func NewTaskService(repo *repositories.TaskRepository) *TaskService {
	return &TaskService{TaskRepo: repo}
}

func (s *TaskService) CreateTask(title, description string, ownerID primitive.ObjectID) error {
	task := &models.Task{
		Title:       title,
		Description: description,
		Status:      models.StatusPending,
		OwnerID:     ownerID,
	}

	return s.TaskRepo.CreateTask(task)
}

func (s *TaskService) GetTasks(ownerID primitive.ObjectID) ([]models.Task, error) {
	return s.TaskRepo.GetTasksByOwner(ownerID)
}
