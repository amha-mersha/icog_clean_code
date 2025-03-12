package usecase

import (
	"strings"
	"time"

	"github.com/amha-mersha/icog_clean_code/internal/domain"
	"github.com/amha-mersha/icog_clean_code/internal/domain/dto"
	"github.com/amha-mersha/icog_clean_code/internal/repository"
	"github.com/google/uuid"
)

type taskItemUC struct {
	repo repository.TaskRepository
}

func NewTaskUseCase(repo repository.TaskRepository) taskItemUC {
	return taskItemUC{repo}
}

func (taskUC *taskItemUC) CreateTask(task *dto.TaskCreateDTO) error {
	// Trim spaces
	task.Title = strings.TrimSpace(task.Title)
	task.Description = strings.TrimSpace(task.Description)

	// Validating task fields
	if task.Title == "" || len(task.Title) > 255 {
		return domain.NewCustomeError(domain.ERR_BAD_REQUEST, "task title can't be more than 255 characters.")
	}
	if len(task.Description) > 1000 {
		return domain.NewCustomeError(domain.ERR_BAD_REQUEST, "task description can't be more than 1000 characters.")
	}
	if task.Deadline.Before(time.Now()) {
		return domain.NewCustomeError(domain.ERR_BAD_REQUEST, "deadline must be a future date.")
	}
	if !dto.ValidStatus(task.Status) {
		task.Status = "pending"
	}
	// Create the task id and pass to the repository
	newTask := domain.TaskItem{
		ID:          uuid.New(),
		Title:       task.Title,
		Description: task.Description,
		CreatedAt:   time.Now(),
		Deadline:    task.Deadline,
		Status:      task.Status,
	}

	return taskUC.repo.Create(&newTask)
}

func (taskUC *taskItemUC) GetTask(id uuid.UUID) (*domain.TaskItem, error) {
	return taskUC.repo.GetByID(id)
}

func (taskUC *taskItemUC) ListTasks() ([]domain.TaskItem, error) {
	return taskUC.repo.GetAll()
}

func (taskUC *taskItemUC) UpdateTask(task *dto.TaskUpdateDTO) error {
	// Trim spaces
	task.Title = strings.TrimSpace(task.Title)
	task.Description = strings.TrimSpace(task.Description)

	// Validating task fields
	if task.Title == "" || len(task.Title) > 255 {
		return domain.NewCustomeError(domain.ERR_BAD_REQUEST, "task title can't be more than 255 characters.")
	}
	if len(task.Description) > 1000 {
		return domain.NewCustomeError(domain.ERR_BAD_REQUEST, "task description can't be more than 1000 characters.")
	}
	if task.Deadline.Before(time.Now()) {
		return domain.NewCustomeError(domain.ERR_BAD_REQUEST, "deadline must be a future date.")
	}
	if !dto.ValidStatus(task.Status) {
		return domain.NewCustomeError(domain.ERR_BAD_REQUEST, "status must be pending, in_progress, or completed.")
	}
	return taskUC.repo.Update(task.ToTaskModel())
}

func (taskUC *taskItemUC) DeleteTask(id uuid.UUID) error {
	return taskUC.repo.Delete(id)
}

func (taskUC *taskItemUC) GetTaskByStatus(status string) ([]domain.TaskItem, error) {
	if !dto.ValidStatus(status) {
		return nil, domain.NewCustomeError(domain.ERR_BAD_REQUEST, "status must be pending, in_progress, or completed.")
	}
	return taskUC.repo.GetByKey("status", status)
}
