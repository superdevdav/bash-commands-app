package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		databaseURL = "user=david password=qwerty123 host=localhost dbname=bash_test"
	}

	os.Exit(m.Run())
}
