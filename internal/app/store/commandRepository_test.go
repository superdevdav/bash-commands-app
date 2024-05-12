package store_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/superdevdav/bash-app/internal/app/model"
	"github.com/superdevdav/bash-app/internal/app/store"
)

// Тестирование функции Create
func TestCommandRepository_Create(t *testing.T) {
	// Подключение к бд
	databaseURL := "user=david password=qwerty123 host=localhost dbname=bash_test"
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("commands")

	comm, err := s.Command().Create(&model.Command{
		Command_name: "ls",
		Result:       "apiserver  cmd  configs  go.mod  go.sum  internal  Makefile  migrations",
	})

	assert.NoError(t, err)
	assert.NotNil(t, comm)
}

// Тестирование функции GetCommandByID
func TestCommandRepository_GetCommandByID(t *testing.T) {
	// Подключение к бд
	databaseURL := "user=david password=qwerty123 host=localhost dbname=bash_test"
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("commands")

	// Несуществующий id
	id := 123
	_, err := s.Command().GetCommandByID(id)
	assert.Error(t, err)
}
