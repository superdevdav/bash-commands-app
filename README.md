## Тестовое задание
Разработать приложение, которое будет предоставлять REST API для запуска команд. Команда представляет собой bash-скрипт. Приложение должно позволять параллельно запускать произвольное количество команд.

## Функционал
1. Создание команды (сохранение ее результата выполнения в БД)
2. Получение списка команд
3. Получение одной команды
4. Удаление команды
5. Тестирование

## Сущность команды (command.go)
У структуры Command имеются следующие поля ID, Command_name, Result, Date_time. Решил добавить Date_time, т.к. это практичнее, например, у нас в БД есть одинаковые команды, следовательно нам нужна дата и время.

## Структура проекта
|--cmd
|   |--apiserver
|       |--main.go
|--configs
|   |--apiserver.toml
|--internal
|   |--app
|       |--apiserver
|       |   |--apiserver_internal_test.go
|       |   |--apiserver.go
|       |   |--config.go
|       |   |--server.go
|       |--model
|       |   |--command_test.go
|       |   |--command.go
|       |   |--testing.go
|       |--store
|           |--sqlstore
|           |   |--commandRepository_test.go
|           |   |--commandRepository.go
|           |   |--store_test.go
|           |   |--store.go
|           |   |--testing.go
|           |--repository.go
|           |--store.go
|--migrations
|   |--20240510000500_create_commands.down.sql
|   |--20240510000500_create_commands.up.sql
|--apiserver
|--go.mod
|--go.sum
|--Makefile

## Про тесты
Запуск тестов происходит командой make test.
Тесты находятся в файлах:
1. apiserver_internal_test.go (тестирование запросов)
2. command_test.go (тестирование структуры Command)
3. commandRepository_test.go (тестирование основного функционала: добавление, получение, удаление)

## Про конфигурационный файл
В apiserver.toml находится конфигурационная информация, такая как номер порта, уровень логгирования, database_url (нужно поменять).

## Запросы
1. Главная страница "/commands"
2. Создание команды "/commands/create?command_name={name}"
3. Получение команды "/commands/get?id={id}"
4. Получение списка команд "/commands/get_all_commands"
5. Удаление команды "/commands/delete?id={id}"

## Запуск
Сервер запускается через файл apiserver (./apiserver)

## Используемые технологии
1. Linux (Ubuntu 24.04 LTS)
2. Golang (Стандартные библиотеки)
3. Postgres 16
