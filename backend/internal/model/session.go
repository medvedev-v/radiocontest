package model

import (
    "time"
)

type Session struct {
    ID           string     `json:"id" db:"id"`                     // внутренний идентификатор (можно использовать UUID)
    UserID       int64      `json:"user_id" db:"user_id"`
    Token        string     `json:"token" db:"token"`               // уникальный токен доступа
    RefreshToken *string    `json:"refresh_token,omitempty" db:"refresh_token"` // опционально
    IPAddress    string     `json:"ip_address" db:"ip_address"`
    UserAgent    string     `json:"user_agent" db:"user_agent"`
    ExpiresAt    time.Time  `json:"expires_at" db:"expires_at"`
    CreatedAt    time.Time  `json:"created_at" db:"created_at"`
    UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
    IsActive     bool       `json:"is_active" db:"is_active"`       // флаг для "мягкого" удаления
}
