package model

import validation "github.com/go-ozzo/ozzo-validation"

// Сущность Command
type Command struct {
	ID           int
	Command_name string
	Result       string
	Date_time    string
}

// Валидация структуры Command
func (c *Command) ValidateCommand() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.Command_name, validation.Required),
		validation.Field(&c.Date_time, validation.Required),
	)
}
