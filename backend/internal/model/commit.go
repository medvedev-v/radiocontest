package model

import (
	"time"
)

type Commit struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	TeamID    int       `json:"team_id"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}
