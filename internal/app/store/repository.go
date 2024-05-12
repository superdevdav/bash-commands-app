package store

import "github.com/superdevdav/bash-app/internal/app/model"

type CommandRepository interface {
	Create(*model.Command) error
	GetCommandByID(int) (*model.Command, error)
	GetAllCommands() ([]*model.Command, error)
	DeleteCommandByID(int) error
}
