package usecase

import (
	"github.com/amha-mersha/icog_clean_code/internal/domain"
	"github.com/amha-mersha/icog_clean_code/internal/domain/dto"
	"github.com/google/uuid"
)

type TaskUsecase interface {
	CreateTask(task *dto.TaskCreateDTO) error
	GetTask(id uuid.UUID) (*domain.TaskItem, error)
	GetTaskByStatus(status string) ([]domain.TaskItem, error)
	ListTasks() ([]domain.TaskItem, error)
	UpdateTask(task *dto.TaskUpdateDTO) error
	DeleteTask(id uuid.UUID) error
}
