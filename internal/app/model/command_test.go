package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/superdevdav/bash-app/internal/app/model"
)

// Тестирование валидации структуры Command
func TestCommand_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		c       func() *model.Command
		isValid bool
	}{
		{
			name: "valid",
			c: func() *model.Command {
				return model.TestCommand(t)
			},
			isValid: true,
		},
		{
			name: "empty Command_name",
			c: func() *model.Command {
				c := model.TestCommand(t)
				c.Command_name = ""
				return c
			},
			isValid: false,
		},
		{
			name: "empty Date_time",
			c: func() *model.Command {
				c := model.TestCommand(t)
				c.Date_time = ""
				return c
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		if tc.isValid {
			assert.NoError(t, tc.c().ValidateCommand())
		} else {
			assert.Error(t, tc.c().ValidateCommand())
		}
	}
}
