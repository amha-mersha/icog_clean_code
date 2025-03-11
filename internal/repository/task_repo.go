package repository

import (
	"errors"
	"net/http"

	"github.com/amha-mersha/icog_clean_code/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type taskItemRepo struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) taskItemRepo {
	return taskItemRepo{
		db,
	}
}

func (repo *taskItemRepo) Create(task *domain.TaskItem) error {
	if result := repo.db.Create(&task); result.Error != nil {
		return domain.NewCustomeError(http.StatusInternalServerError, result.Error.Error())
	}
	return nil
}

func (repo *taskItemRepo) GetByID(id uuid.UUID) (*domain.TaskItem, error) {
	var existingTask domain.TaskItem
	if err := repo.db.Take(&existingTask, id); err.Error != nil {
		return nil, domain.NewCustomeError(http.StatusInternalServerError, err.Error.Error())
	}
	return &existingTask, nil
}

func (repo *taskItemRepo) GetAll() ([]domain.TaskItem, error) {
	var existingTasks []domain.TaskItem
	if err := repo.db.Find(&existingTasks); err.Error != nil {
		return nil, domain.NewCustomeError(http.StatusInternalServerError, err.Error.Error())
	}
	return existingTasks, nil
}

func (repo *taskItemRepo) Update(task *domain.TaskItem) error {
	if err := repo.db.First(&task, task.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.NewCustomeError(http.StatusNotFound, "Task not found")
		}
		return domain.NewCustomeError(http.StatusInternalServerError, err.Error())
	}

	if err := repo.db.Save(&task).Error; err != nil {
		return domain.NewCustomeError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (repo *taskItemRepo) Delete(id uuid.UUID) error {
	var task domain.TaskItem
	if err := repo.db.First(&task, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.NewCustomeError(http.StatusNotFound, "Task not found")
		}
		return domain.NewCustomeError(http.StatusInternalServerError, err.Error())
	}

	if err := repo.db.Delete(&task).Error; err != nil {
		return domain.NewCustomeError(http.StatusInternalServerError, err.Error())
	}

	return nil
}
