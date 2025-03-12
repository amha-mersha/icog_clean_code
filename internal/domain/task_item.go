package domain

import (
	"time"

	"github.com/google/uuid"
)

type TaskItem struct {
	ID          uuid.UUID `json:"id" example:"f3f905c2-f40c-441b-8e74-f84923b3d158"`
	CreatedAt   time.Time `json:"created_at"`
	Title       string    `json:"title" example:"Do the dishes"`
	Description string    `json:"description" example:"Load the dishwasher with all the plates"`
	Deadline    time.Time `json:"deadline"`
	Status      string    `json:"status" example:"pending"`
}
