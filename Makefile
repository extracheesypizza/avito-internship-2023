include .env
export

build:
	docker-compose build avito-app

run:
	docker-compose up avito-app