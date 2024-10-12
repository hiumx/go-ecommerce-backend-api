# name app

APP_NAME = server

dev:
	go run ./cmd/${APP_NAME}

run:
	docker compose up -d && go run ./cmd/${APP_NAME}/

up:
	docker compose up -d

up_build:
	docker compose up -d --build

down:
	docker compose down

kill:
	docker compose kill