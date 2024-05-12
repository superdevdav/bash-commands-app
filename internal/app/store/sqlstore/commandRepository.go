package sqlstore

import (
	"github.com/superdevdav/bash-app/internal/app/model"
)

type CommandRepository struct {
	store *Store
}

// Создание команды
func (r *CommandRepository) Create(comm *model.Command) error {
	// Проверка на валидацию
	if err := comm.ValidateCommand(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO commands (command_name, result, date_time) VALUES ($1, $2, $3) RETURNING id",
		comm.Command_name, comm.Result, comm.Date_time,
	).Scan(&comm.ID)
}

// Получение команды по id
func (r *CommandRepository) GetCommandByID(id int) (*model.Command, error) {
	comm := &model.Command{}
	if err := r.store.db.QueryRow(
		"SELECT id, command_name, result, date_time FROM commands WHERE id = $1;", id,
	).Scan(&comm.ID, &comm.Command_name, &comm.Result, &comm.Date_time); err != nil {
		return nil, err
	}

	// Проверка на валидацию команды из бд
	if err := comm.ValidateCommand(); err != nil {
		return nil, err
	}

	return comm, nil
}

// Получение всех команд
func (r *CommandRepository) GetAllCommands() ([]*model.Command, error) {
	rows, err := r.store.db.Query("SELECT id, command_name, result, date_time FROM commands")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Создание массива для хранения команд
	commands := []*model.Command{}

	// Сканирование результатов запроса
	for rows.Next() {
		comm := &model.Command{}
		if err := rows.Scan(&comm.ID, &comm.Command_name, &comm.Result, &comm.Date_time); err != nil {
			return nil, err
		}

		// Проверка на валидацию
		if err := comm.ValidateCommand(); err != nil {
			continue
		}

		commands = append(commands, comm)
	}

	// Проверка на наличие ошибок после rows.Scan
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return commands, nil
}

// Удаление команды по id
func (r *CommandRepository) DeleteCommandByID(id int) error {
	_, err := r.store.db.Exec("DELETE FROM commands WHERE id = $1;", id)
	if err != nil {
		return err
	}
	return nil
}
