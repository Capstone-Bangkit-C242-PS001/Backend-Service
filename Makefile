ifneq (,$(wildcard .env))
    include .env
    export
endif

start:
	docker-compose up -d

start-build:
	docker-compose up --build -d

stop:
	docker-compose down 

restart: stop start

create-migration: 
ifeq ($(OS),Windows_NT)
	@if "$(name)" == "" ( \
		echo "Error: name parameter is required. Usage: make create-migration name=<migration_name>";\
		exit 1; \
	) else ( \
		migrate create -ext sql -dir database/migration -seq $(name) \
	)
else
	@if [ -z "$(name)" ]; then \
		echo "Error: name parameter is required. Usage: make create-migration name=<migration_name>"; \
		exit 1; \
	fi; \
	migrate create -ext sql -dir database/migration -seq $(name)
endif

debug-env:
	@echo "DB_USER: '$(DB_USER)'"
	@echo "DB_PASSWORD: '$(DB_PASSWORD)'"
	@echo "DB_HOST: '$(DB_HOST)'"
	@echo "DB_PORT: '$(DB_PORT)'"
	@echo "DB_NAME: '$(DB_NAME)'"

migration-up:
	@echo "Running migrations up..."
	migrate -path database/migration -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp(localhost:$(DB_PORT))/$(DB_NAME)" up

migration-down:
	@echo "Running migrations down..."
	migrate -path database/migration -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp(localhost:$(DB_PORT))/$(DB_NAME)" down 1

migration-force:
ifeq ($(OS),Windows_NT)
	@if "$(version)" == "" ( \
		echo "Error: version parameter is required. Usage: make migration-force version=<migration_version>" && exit 1 \
	) else ( \
		echo "Force migrations to $(version)" && \
		migrate -path database/migration -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp(localhost:$(DB_PORT))/$(DB_NAME)" force $(version) \
	)
else
	@if [ -z "$(version)" ]; then \
		echo "Error: version parameter is required. Usage: make migration-force version=<migration_version>"; \
		exit 1; \
	fi; \
	echo "Force migrations to $(version)"; \
	migrate -path database/migration -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp(localhost:$(DB_PORT))/$(DB_NAME)" force $(version)
endif

swag:
	swag init --generalInfo cmd/main.go --output docs
