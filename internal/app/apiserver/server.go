package apiserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/superdevdav/bash-app/internal/app/model"
	"github.com/superdevdav/bash-app/internal/app/store"
)

// server ...
type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

// newServer ...
func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	s.configureRouter()

	return s
}

// ServerHTTP ...
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// ConfigureRouter ...
func (s *server) configureRouter() {
	s.router.HandleFunc("/commands", s.handleIndex()).Methods("GET")
	s.router.HandleFunc("/commands/create", s.handleCreate())
	s.router.HandleFunc("/commands/get", s.handleGetCommandByID())
	s.router.HandleFunc("/commands/get_all_commands", s.handleGetAllCommands())
	s.router.HandleFunc("/commands/delete", s.handleDeleteCommandByID())
}

// Главная страница
func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Сервис по работе с bash командами")
	}
}

// Создание команды
func (s *server) handleCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		command := r.URL.Query().Get("command_name") // Команда из запроса

		currentTime := time.Now() // Текущее время UTC
		formattedDateTime := currentTime.Format("02.01.2006 15:04:05")

		// Создание команды
		cmd := exec.Command("bash", "-c", command)

		// Запуска команды и захват вывода
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()

		// Проверка на корректность команды
		if err == nil {

			// Создание структуры Command
			newCommand := &model.Command{
				Command_name: command,
				Result:       out.String(),
				Date_time:    formattedDateTime,
			}

			err := s.store.Command().Create(newCommand)

			if err != nil {
				http.Error(w, "Failed to create command", http.StatusInternalServerError)
				s.logger.Error(fmt.Sprintf("Failed to create command: %s", command))
				return
			}

			s.logger.Info(fmt.Sprintf("Command %s added successfully", command))

			io.WriteString(w, fmt.Sprintf("Команда %s успешно добавлена", command))

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w)
		} else {
			io.WriteString(w, "Некорректная команда")

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w)
		}
	}
}

// Получение команды по id
func (s *server) handleGetCommandByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("ConfigureRouterGetByID is OK")
		idStr := r.URL.Query().Get("id") // id из запроса
		id, _ := strconv.Atoi(idStr)

		command, err := s.store.Command().GetCommandByID(id)
		if err != nil {
			s.logger.Error(err)
			http.Error(w, "Failed to get command", http.StatusInternalServerError)
			return
		}

		s.logger.Info(fmt.Sprintf("Command with id = %d getted successfully", command.ID))

		// Преобразование команды string -> json
		data, err := json.Marshal(command)
		if err != nil {
			s.logger.Error(err)
			http.Error(w, "Failed to encode command to JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		// Запись данных json в ResponseWriter
		_, err = w.Write(data)
		if err != nil {
			s.logger.Error(err)
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
	}
}

// Получение всех команд
func (s *server) handleGetAllCommands() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("ConfigureRouterGetAllCommands is OK")
		commands, err := s.store.Command().GetAllCommands()

		if err != nil {
			s.logger.Error(err)
			http.Error(w, "Failed to get all commands", http.StatusInternalServerError)
			return
		}

		s.logger.Info("All commands get OK")

		// Преобразование массива команд в JSON
		data, err := json.Marshal(commands)
		if err != nil {
			s.logger.Error(err)
			http.Error(w, "Failed to encode commands to JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		// Запись JSON-данных в ResponseWriter
		_, err = w.Write(data)
		if err != nil {
			s.logger.Error(err)
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
	}
}

// Удаление команды по id
func (s *server) handleDeleteCommandByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id") // id из запроса
		id, _ := strconv.Atoi(idStr)

		err := s.store.Command().DeleteCommandByID(id)
		if err != nil {
			s.logger.Error(err)
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}

		s.logger.Info(fmt.Sprintf("Command with id = %d deleted", id))
		io.WriteString(w, fmt.Sprintf("Команда с id = %d удалена", id))

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w)
	}
}
