package model

import (
	"time"
)

type TeamMember struct {
	ID     int `json:"id"`
	TeamID int `json:"team_id"`
	UserID int `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
