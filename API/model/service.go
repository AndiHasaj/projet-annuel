package model

import (
	"time"
)

type Service struct {
	ID          int       `json:"id"`
	TypeService string    `json:"type_service"`
	Titre       string    `json:"titre"`
	Description string    `json:"description"`
	Prix        float32   `json:"prix"`
	DateService time.Time `json:"date_service"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
