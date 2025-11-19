# Makefile for TodoBackend

# Переменные для миграций
DB_DSN := "postgres://postgres:yourpassword@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Переменные для приложения
APP_PATH := cmd/main.go

.PHONY: run migrate migrate-down migrate-new

# Запуск приложения
run:
	go run $(APP_PATH)

# Миграции ==========================================

# Создание новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations $(NAME)

# Применение миграций
migrate:
	$(MIGRATE) up

# Откат миграций (с автоматическим подтверждением)
migrate-down:
	echo "y" | $(MIGRATE) down

# Просмотр статуса миграций
migrate-status:
	$(MIGRATE) version

# Очистка БД (острожно - удаляет таблицу!)
db-clean:
	psql -h localhost -U postgres -d postgres -c "DROP TABLE IF EXISTS tasks;"