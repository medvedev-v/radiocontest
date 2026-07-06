package model

import (
	"time"
)

type Teammate struct {
	ID     int `json:"id"`
	TeamID int `json:"team_id"`
	UserID int `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
