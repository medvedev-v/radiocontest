package repository

import (
	"database/sql"
	"fmt"
	model "github.com/medvedev-v/radiocontest/internal/model"
)

// UserRepository представляет репозиторий для работы с пользователями
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository создает новый экземпляр UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create создает нового пользователя
func (r *UserRepository) Create(user model.User) (int, error) {
	query := `INSERT INTO users (username, email, passwordhash, role, createdat) VALUES (?, ?, ?, ?, ?)`
	result, err := r.db.Exec(query, user.UserName, user.Email, user.PasswordHash, user.UserRole, user.CreatedAt)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert ID: %w", err)
	}
	return int(id), nil
}

// GetByID возвращает пользователя по ID
func (r *UserRepository) GetByID(id int) (*model.User, error) {
	query := `SELECT id, username, email, role, created_at FROM users WHERE id = ?`
	row := r.db.QueryRow(query, id)

	var user model.User
	err := row.Scan(&user.ID, &user.UserName, &user.Email, &user.UserRole, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found with ID: %d", id)
		}
		return nil, fmt.Errorf("failed to scan user: %w", err)
	}
	return &user, nil
}

// Update обновляет данные пользователя
func (r *UserRepository) Update(user model.User) error {
	query := `UPDATE users SET username = ?, email = ?, role = ?, createdat = ? WHERE id = ?`
	_, err := r.db.Exec(query, user.UserName, user.Email, user.UserRole, user.CreatedAt, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

// Delete удаляет пользователя по ID
func (r *UserRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
