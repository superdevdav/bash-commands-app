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

type APIserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// Создание apiserver
func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Запуск сервера
func (s *APIserver) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()
	s.configureRouterCreate()
	s.configureRouterGetByID()
	s.configureRouterGetAllCommands()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("Starting API server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// Конфигурация БД
func (s *APIserver) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}

// Функция для конфигурации логера
func (s *APIserver) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.logLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

// Роутинг (главная страница)
func (s *APIserver) configureRouter() {
	s.router.HandleFunc("/commands", s.handleIndex())
}

func (s *APIserver) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Сервис по работе с bash командами")
	}
}

// Роутинг (добавление команды)
func (s *APIserver) configureRouterCreate() {
	s.router.HandleFunc("/commands/create/{command_name}", s.handleCreate())
}

func (s *APIserver) handleCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		command := mux.Vars(r)["command_name"] // Команда из запроса

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

			_, err := s.store.Command().Create(newCommand)

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

// Роутинг (получение команды по id)
func (s *APIserver) configureRouterGetByID() {
	s.router.HandleFunc("/commands/get/{id}", s.handleGetByID())
}

func (s *APIserver) handleGetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("ConfigureRouterGetByID is OK")
		idStr := mux.Vars(r)["id"] // id из запроса
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

// Роутинг (получение всех команд)
func (s *APIserver) configureRouterGetAllCommands() {
	s.router.HandleFunc("/commands/get", s.handleGetAllCommands())
}

func (s *APIserver) handleGetAllCommands() http.HandlerFunc {
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
