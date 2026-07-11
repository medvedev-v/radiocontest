package repository

import (
	"database/sql"
	"fmt"
	model "github.com/medvedev-v/radiocontest/internal/model"
)

type OperatorRepository struct {
	db *sql.DB
}

func NewOperatorRepository(db *sql.DB) *OperatorRepository {
	return &OperatorRepository{db: db}
}

func (r *OperatorRepository) Create(operator model.Operator) (int, error) {
	query := `INSERT INTO operators (callsign) VALUES (?)`
	result, err := r.db.Exec(query, operator.Callsign)
	if err != nil {
		return 0, fmt.Errorf("failed to create operator: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert ID: %w", err)
	}
	return int(id), nil
}

func (r *OperatorRepository) GetByID(id int) (*model.Operator, error) {
	query := `SELECT id, callsign FROM operators WHERE id = ?`
	row := r.db.QueryRow(query, id)
	var operator model.Operator
	err := row.Scan(&operator.ID, &operator.Callsign)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("operator not found with ID: %d", id)
		}
		return nil, fmt.Errorf("failed to scan operator: %w", err)
	}
	return &operator, nil
}

func (r *OperatorRepository) Update(operator model.Operator) error {
	query := `UPDATE operators SET callsign = ? WHERE id = ?`
	_, err := r.db.Exec(query, operator.Callsign, operator.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func (r *OperatorRepository) Delete(id int) error {
	query := `DELETE FROM operators WHERE id = ?`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete operator: %w", err)
	}
	return nil
}
