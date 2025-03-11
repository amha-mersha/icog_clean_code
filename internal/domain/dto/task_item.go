package dto

import (
	"time"
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

func ValidStatus(status string) bool {
	return !validStatuses[status]
}
