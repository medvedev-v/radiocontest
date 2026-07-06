package migrations

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func RunMigrations(db *sql.DB, migrationsPath string) error {
	files, err := os.ReadDir(migrationsPath)
	if err != nil {
		return fmt.Errorf("failed to read migrations directory: %v", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" {
			migrationPath := filepath.Join(migrationsPath, file.Name())
			migrationSQL, err := os.ReadFile(migrationPath)
			if err != nil {
				return fmt.Errorf("failed to read migration file %s: %v", file.Name(), err)
			}

			_, err = db.Exec(string(migrationSQL))
			if err != nil {
				return fmt.Errorf("failed to execute migration %s: %v", file.Name(), err)
			}

			log.Printf("Migration %s applied successfully", file.Name())
		}
	}

	return nil
}
