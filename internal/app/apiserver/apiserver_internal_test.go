package apiserver

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/superdevdav/bash-app/internal/app/model"
	"github.com/superdevdav/bash-app/internal/app/store/sqlstore"
)

var (
	databaseURL = "user=david password=qwerty123 host=localhost dbname=bash_test" // тестовая БД
)

// Тестирование HandleIndex (Главная страница)
func TestServer_HandleIndex(t *testing.T) {
	// Подключение к БД
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("commands")

	store := sqlstore.New(db)

	srv := newServer(store)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/commands", nil)
	srv.handleIndex().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Сервис по работе с bash командами", http.StatusOK)
}

// Тестирование HandleCreate (некорректная команда)
func TestServer_HandleCreate_invalidCommand(t *testing.T) {
	// Подключение к БД
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("commands")

	store := sqlstore.New(db)

	srv := newServer(store)
	rec := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodPost, "/commands/create?command_name=incorrect-command", nil)
	if err != nil {
		t.Fatal(err)
	}
	srv.handleCreate().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Некорректная команда")
}

// Тестирование HandleCreate (корректная команда)
func TestServer_HandleCreate_validCommand(t *testing.T) {
	// Подключение к БД
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("commands")

	store := sqlstore.New(db)

	srv := newServer(store)
	rec := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodPost, "/commands/create?command_name=pwd", nil)
	if err != nil {
		t.Fatal(err)
	}
	srv.handleCreate().ServeHTTP(rec, req)
	assert.Equal(t, "Команда pwd успешно добавлена", rec.Body.String())
}

// Тестирование HandleGetCommandByID
func TestServer_HandleGetCommandByID_correctCommand(t *testing.T) {
	// Подключение к БД
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("commands")

	store := sqlstore.New(db)

	srv := newServer(store)
	rec := httptest.NewRecorder()

	// Тестовая команда
	testCommand := *model.TestCommand(t)
	store.Command().Create(&testCommand)
	id := testCommand.ID

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/commands/get?id=%d", id), nil)
	if err != nil {
		t.Fatal(err)
	}

	srv.handleGetCommandByID().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status OK; got %v", rec.Code)
	}
}

// Тестирование handleDeleteCommandByID
func TestServer_handleDeleteCommandByID(t *testing.T) {
	// Подключение к БД
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("commands")

	store := sqlstore.New(db)

	srv := newServer(store)
	rec := httptest.NewRecorder()

	// Тестовая команда
	testCommand := *model.TestCommand(t)
	store.Command().Create(&testCommand)
	id := testCommand.ID

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/commands/delete?id=%d", id), nil)
	if err != nil {
		t.Fatal(err)
	}
	srv.handleDeleteCommandByID().ServeHTTP(rec, req)
	assert.Equal(t, fmt.Sprintf("Команда с id = %d удалена", id), rec.Body.String())
}
