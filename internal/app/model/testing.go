package model

import "testing"

// Тестовая команда
func TestCommand(t *testing.T) *Command {
	t.Helper()

	return &Command{
		ID:           123,
		Command_name: "ls",
		Result:       "apiserver  cmd  configs  go.mod  go.sum  internal  Makefile  migrations",
		Date_time:    "02.01.2006 15:04:05",
	}
}
