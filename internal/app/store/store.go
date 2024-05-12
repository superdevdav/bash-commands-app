package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Структура хранилища
type Store struct {
	config            *Config
	db                *sql.DB
	commandRepository *CommandRepository
}

// Новое хранилище
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)

	if err != nil {
		return err
	}

	// Проверка соединения
	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

// Close ...
func (s *Store) Close() {
	s.db.Close()
}

// Чтобы "внешний мир" мог использовать репозиторий
func (s *Store) Command() *CommandRepository {
	if s.commandRepository != nil {
		return s.commandRepository
	}

	// Если нету репозитория, то создаем его
	s.commandRepository = &CommandRepository{
		store: s,
	}

	return s.commandRepository
}
