package sqlstore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/superdevdav/bash-app/internal/app/model"
	"github.com/superdevdav/bash-app/internal/app/store/sqlstore"
)

var (
	databaseURL = "user=david password=qwerty123 host=localhost dbname=bash_test" // тестовая БД
)

// Тестирование функции Create
func TestCommandRepository_Create(t *testing.T) {
	// Подключение к БД
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("commands")

	s := sqlstore.New(db)

	err := s.Command().Create(model.TestCommand(t))

	assert.NoError(t, err)
}

// Тестирование функции GetCommandByID (НЕСУЩЕСТВУЮЩАЯ КОМАНДА)
func TestCommandRepository_GetCommandByIDNotExisted(t *testing.T) {
	// Подключение к БД
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("commands")

	s := sqlstore.New(db)

	// Несуществующий id
	id := 123
	_, err := s.Command().GetCommandByID(id)
	assert.Error(t, err)
}

// Тестирование функции GetCommandByID (СУЩЕСТВУЮЩАЯ КОМАНДА)
func TestCommandRepository_GetCommandByIDExisted(t *testing.T) {
	// Подключение к БД
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("commands")

	s := sqlstore.New(db)

	// Тестовая команда
	testCommand := *model.TestCommand(t)
	s.Command().Create(&testCommand)
	id := testCommand.ID

	c, err := s.Command().GetCommandByID(id)

	assert.NoError(t, err)
	assert.NotNil(t, c)
}

// Тестирование функции GetAllCommands
func TestCommandRepository_GetAllCommands(t *testing.T) {
	// Подключение к БД
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("commands")

	s := sqlstore.New(db)

	command1 := model.Command{
		Command_name: "name1",
		Result:       "result1",
		Date_time:    "date_time1",
	}
	s.Command().Create(&command1)

	command2 := model.Command{
		Command_name: "name2",
		Result:       "result2",
		Date_time:    "date_time2",
	}
	s.Command().Create(&command2)

	command3 := model.Command{
		Command_name: "name3",
		Result:       "result3",
		Date_time:    "date_time3",
	}
	s.Command().Create(&command3)

	commands, err := s.Command().GetAllCommands()
	if err != nil {
		assert.Error(t, err)
	}

	assert.NotNil(t, commands)
}

// Тестирование фукнции DeleteCommandByID
func TestCommandRepository_Delete(t *testing.T) {
	// Подключение к БД
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("commands")

	s := sqlstore.New(db)

	// Тестовая команда
	testCommand := *model.TestCommand(t)
	id := testCommand.ID

	s.Command().Create(&testCommand)

	err := s.Command().DeleteCommandByID(id)

	assert.NoError(t, err)
}
