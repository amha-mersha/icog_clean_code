package repository

import (
	"errors"
	"fmt"

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
		return domain.NewCustomeError(domain.ERR_INTERNAL_SERVER, result.Error.Error())
	}
	return nil
}

func (repo *taskItemRepo) GetByID(id uuid.UUID) (*domain.TaskItem, error) {
	var existingTask domain.TaskItem
	if err := repo.db.Take(&existingTask, id); err.Error != nil {
		return nil, domain.NewCustomeError(domain.ERR_INTERNAL_SERVER, err.Error.Error())
	}
	return &existingTask, nil
}

func (repo *taskItemRepo) GetAll() ([]domain.TaskItem, error) {
	var existingTasks []domain.TaskItem
	if err := repo.db.Find(&existingTasks); err.Error != nil {
		return nil, domain.NewCustomeError(domain.ERR_INTERNAL_SERVER, err.Error.Error())
	}
	return existingTasks, nil
}

func (repo *taskItemRepo) Update(task *domain.TaskItem) error {
	if err := repo.db.First(&task, task.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.NewCustomeError(domain.ERR_NOT_FOUND, "Task not found")
		}
		return domain.NewCustomeError(domain.ERR_INTERNAL_SERVER, err.Error())
	}

	if err := repo.db.Model(&task).Updates(domain.TaskItem{
		Title:       task.Title,
		Description: task.Description,
		Deadline:    task.Deadline,
		Status:      task.Status,
	}).Error; err != nil {
		return domain.NewCustomeError(domain.ERR_INTERNAL_SERVER, err.Error())
	}

	return nil
}

func (repo *taskItemRepo) Delete(id uuid.UUID) error {
	var task domain.TaskItem
	if err := repo.db.First(&task, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.NewCustomeError(domain.ERR_NOT_FOUND, "Task not found")
		}
		return domain.NewCustomeError(domain.ERR_INTERNAL_SERVER, err.Error())
	}

	if err := repo.db.Delete(&task).Error; err != nil {
		return domain.NewCustomeError(domain.ERR_INTERNAL_SERVER, err.Error())
	}

	return nil
}

func (repo *taskItemRepo) GetByKey(key string, value interface{}) ([]domain.TaskItem, error) {
	var tasks []domain.TaskItem
	if err := repo.db.Where(fmt.Sprintf("%s=?", key), value).Find(&tasks).Error; err != nil {
		return nil, domain.NewCustomeError(domain.ERR_INTERNAL_SERVER, err.Error())
	}
	return tasks, nil
}
