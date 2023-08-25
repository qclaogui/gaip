package todo

import (
	"github.com/google/uuid"
	"time"
)

type Todo struct {
	ID          uuid.UUID
	Title       string
	Description string
	Reminder    time.Time
}
