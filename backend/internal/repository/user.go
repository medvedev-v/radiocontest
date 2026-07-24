package repository

import (
	"database/sql"
	"fmt"

	model "github.com/medvedev-v/radiocontest/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create добавляет нового пользователя.
// Возвращает ID созданной записи.
func (r *UserRepository) Create(user model.User) (int64, error) {
	query := `INSERT INTO users (callsign, password_hash) VALUES (?, ?)`
	result, err := r.db.Exec(query, user.Callsign, user.PasswordHash)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert ID: %w", err)
	}
	return id, nil
}

// GetByID возвращает пользователя по ID.
func (r *UserRepository) GetByID(id int64) (*model.User, error) {
	query := `SELECT id, callsign, password_hash FROM users WHERE id = ?`
	var user model.User
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Callsign, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found with ID: %d", id)
		}
		return nil, fmt.Errorf("failed to scan user: %w", err)
	}
	return &user, nil
}

// FindByCallsign ищет пользователя по позывному.
func (r *UserRepository) FindByCallsign(callsign string) (*model.User, error) {
	query := `SELECT id, callsign, password_hash FROM users WHERE callsign = ?`
	var user model.User
	err := r.db.QueryRow(query, callsign).Scan(&user.ID, &user.Callsign, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found with callsign: %s", callsign)
		}
		return nil, fmt.Errorf("failed to scan user: %w", err)
	}
	return &user, nil
}

// Update обновляет данные пользователя (callsign и/или password_hash).
func (r *UserRepository) Update(user model.User) error {
	query := `UPDATE users SET callsign = ?, password_hash = ? WHERE id = ?`
	_, err := r.db.Exec(query, user.Callsign, user.PasswordHash, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

// Delete удаляет пользователя по ID.
func (r *UserRepository) Delete(id int64) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
