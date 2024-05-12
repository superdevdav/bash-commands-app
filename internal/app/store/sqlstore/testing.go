package sqlstore

import (
	"database/sql"
	"strings"
	"testing"
)

// Тестовая БД
func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("postgres", databaseURL)

	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			query := "TRUNCATE TABLE " + strings.Join(tables, ", ")
			_, err := db.Exec(query)
			if err != nil {
				t.Fatalf("Error truncating tables: %v", err)
			}
		}
		db.Close()
	}
}
