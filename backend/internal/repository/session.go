package repository

import (
	"database/sql"
	"fmt"
	"time"

	model "github.com/medvedev-v/radiocontest/internal/model"
)

type SessionRepository struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

// Create создаёт новую сессию.
func (r *SessionRepository) Create(session model.Session) error {
	query := `
		INSERT INTO sessions 
		(id, user_id, token, refresh_token, ip_address, user_agent, expires_at, created_at, updated_at, is_active)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	now := time.Now()
	var refreshToken interface{} = nil
	if session.RefreshToken != nil {
		refreshToken = *session.RefreshToken
	}
	_, err := r.db.Exec(
		query,
		session.ID,
		session.UserID,
		session.Token,
		refreshToken,
		session.IPAddress,
		session.UserAgent,
		session.ExpiresAt,
		now,
		now,
		true,
	)
	if err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}
	return nil
}

// GetActiveByToken возвращает активную сессию по токену.
// Проверяет, что сессия активна (is_active = 1) и не истекла (expires_at > NOW()).
func (r *SessionRepository) GetActiveByToken(token string) (*model.Session, error) {
	var session model.Session
	var refreshToken sql.NullString
	query := `
		SELECT id, user_id, token, refresh_token, ip_address, user_agent, expires_at, created_at, updated_at, is_active
		FROM sessions
		WHERE token = ? AND is_active = 1 AND expires_at > NOW()
	`
	err := r.db.QueryRow(query, token).Scan(
		&session.ID,
		&session.UserID,
		&session.Token,
		&refreshToken,
		&session.IPAddress,
		&session.UserAgent,
		&session.ExpiresAt,
		&session.CreatedAt,
		&session.UpdatedAt,
		&session.IsActive,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("session not found or expired with token: %s", token)
		}
		return nil, fmt.Errorf("failed to scan session: %w", err)
	}
	if refreshToken.Valid {
		session.RefreshToken = &refreshToken.String
	}
	return &session, nil
}

// DeactivateByToken деактивирует сессию (логаут) по токену.
func (r *SessionRepository) DeactivateByToken(token string) error {
	query := `UPDATE sessions SET is_active = 0, updated_at = ? WHERE token = ? AND is_active = 1`
	_, err := r.db.Exec(query, time.Now(), token)
	if err != nil {
		return fmt.Errorf("failed to deactivate session: %w", err)
	}
	return nil
}

// DeleteExpired удаляет все истекшие или неактивные сессии (для фоновой очистки).
func (r *SessionRepository) DeleteExpired() error {
	query := `DELETE FROM sessions WHERE expires_at < NOW() OR is_active = 0`
	_, err := r.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to delete expired sessions: %w", err)
	}
	return nil
}
