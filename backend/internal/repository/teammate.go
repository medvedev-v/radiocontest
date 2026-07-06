package repository
/**
import (
	"database/sql"
	"fmt"
	model "github.com/medvedev-v/radiocontest/internal/model"
)

// UserRepository представляет репозиторий для работы с пользователями
type TeammateRepository struct {
	db *sql.DB
}

// NewUserRepository создает новый экземпляр UserRepository
func NewTeammateRepository(db *sql.DB) *TeammateRepository {
	return &TeammateRepository{db: db}
}

// Create создает нового пользователя
func (r *TeammateRepository) Create(teammate model.Teammate) (int, error) {
	query := `INSERT INTO teammates (name, email, passwordhash, role, createdat) VALUES (?, ?, ?, ?, ?)`
	result, err := r.db.Exec(query, teammate.UserName, teammate.Email, teammate.PasswordHash, teammate.UserRole, teammate.CreatedAt)
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
func (r *TeammateRepository) GetByID(id int) (*model.Teammate, error) {
	query := `SELECT id, name, email, role, created_at FROM users WHERE id = ?`
	row := r.db.QueryRow(query, id)

	var user model.Teammate
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
func (r *TeammateRepository) Update(user model.Teammate) error {
	query := `UPDATE users SET name = ?, email = ?, role = ?, createdat = ? WHERE id = ?`
	_, err := r.db.Exec(query, user.UserName, user.Email, user.UserRole, user.CreatedAt, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

// Delete удаляет пользователя по ID
func (r *TeammateRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
**/
