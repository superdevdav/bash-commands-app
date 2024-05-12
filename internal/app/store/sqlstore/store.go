package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/superdevdav/bash-app/internal/app/store"
)

// Структура хранилища
type Store struct {
	db                *sql.DB
	commandRepository *CommandRepository
}

// Новое хранилище
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// Чтобы "внешний мир" мог использовать репозиторий
func (s *Store) Command() store.CommandRepository {
	if s.commandRepository != nil {
		return s.commandRepository
	}

	// Если нету репозитория, то создаем его
	s.commandRepository = &CommandRepository{
		store: s,
	}

	return s.commandRepository
}
