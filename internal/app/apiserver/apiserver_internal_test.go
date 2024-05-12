package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Главная страница
func TestAPIserver_HandleIndex(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/commands", nil)
	s.handleIndex().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Сервис по работе с bash командами")
}

// Добавление команды
func TestAPIserver_HandleCreate(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/commands/create/ls", nil)
	if err != nil {
		t.Fatal(err)
	}
	s.handleCreate().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Команда ls успешно добавлена")
}
