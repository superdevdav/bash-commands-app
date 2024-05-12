package store

type Store interface {
	Command() CommandRepository
}
