package dto

import (
	"time"

	"github.com/amha-mersha/icog_clean_code/internal/domain"
	"github.com/google/uuid"
)

var validStatuses = map[string]bool{
	"pending":     true,
	"in_progress": true,
	"completed":   true,
}

type TaskCreateDTO struct {
	Title       string    `json:"title" binding:"required,max=255" example:"Do the dishes"`
	Description string    `json:"description" binding:"max=1000" example:"Load the dishwasher with all the plates"`
	Deadline    time.Time `json:"deadline" binding:"required"`
	Status      string    `json:"status" binding:"omitempty,oneof=pending in_progress completed" example:"pending"`
}

type TaskUpdateDTO struct {
	ID          uuid.UUID `json:"id" binding:"required" example:"05b2d3b8-8dde-4dc8-85ed-3bc52b7aa3a7"`
	Title       string    `json:"title" binding:"required,max=255" example:"Do the dishes"`
	Description string    `json:"description" binding:"max=1000" example:"Load the dishwasher with all the plates"`
	Deadline    time.Time `json:"deadline" binding:"required"`
	Status      string    `json:"status" binding:"omitempty,oneof=pending in_progress completed" example:"pending"`
}

func (dto *TaskUpdateDTO) ToTaskModel() *domain.TaskItem {
	return &domain.TaskItem{
		ID:          dto.ID,
		Title:       dto.Title,
		Description: dto.Description,
		Status:      dto.Status,
		Deadline:    dto.Deadline,
	}
}

func ValidStatus(status string) bool {
	return validStatuses[status]
}
