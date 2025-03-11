package repository

import (
	"github.com/amha-mersha/icog_clean_code/internal/domain"
	"github.com/amha-mersha/icog_clean_code/internal/domain/dto"
	"github.com/google/uuid"
)

type TaskRepository interface {
	Create(task *domain.TaskItem) error
	GetByID(id uuid.UUID) (*domain.TaskItem, error)
	GetAll() ([]domain.TaskItem, error)
	GetByKey(key string, value interface{}) ([]domain.TaskItem, error)
	Update(task *dto.TaskUpdateDTO) error
	Delete(id uuid.UUID) error
}
