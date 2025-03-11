package usecase

import (
	"github.com/amha-mersha/icog_clean_code/internal/domain"
	"github.com/amha-mersha/icog_clean_code/internal/domain/dto"
	"github.com/google/uuid"
)

type TaskUsecase interface {
	CreateTask(task *dto.TaskCreateDTO) *domain.CustomeError
	GetTask(id uuid.UUID) (*domain.TaskItem, *domain.CustomeError)
	ListTasks() ([]domain.TaskItem, *domain.CustomeError)
	UpdateTask(task *domain.TaskItem) *domain.CustomeError
	DeleteTask(id uuid.UUID) *domain.CustomeError
}
