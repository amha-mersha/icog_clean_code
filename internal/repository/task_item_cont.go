package repository

import (
	"github.com/amha-mersha/icog_clean_code/internal/domain"
	"github.com/google/uuid"
)

type TaskRepository interface {
	Create(task *domain.TaskItem) error
	GetByID(id uuid.UUID) (*domain.TaskItem, error)
	GetAll() ([]domain.TaskItem, error)
	Update(task *domain.TaskItem) error
	Delete(id uuid.UUID) error
}
