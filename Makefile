build_db:
	# Stop the current container
	docker-compose stop postgres-database
	# Build the database application
	docker-compose build --no-cache postgres-database
	# Bring the containers up
	docker-compose up -d postgres-database

build_app:
	# Stop the current container
	docker-compose stop go-url-shortner-application
	# Build the golang application
	docker-compose build --no-cache go-url-shortner-application
	# Bring the containers up
	docker-compose up -d go-url-shortner-application

build_all: build_db build_app

start_all:
	docker-compose up -d

attach_to_db:
	docker exec -it service-postgres-database bash

attach_to_app:
	docker exec -it go-url-shortner-application sh

stop_all:
	docker-compose down

stop_db:
	docker-compose down postgres-database

stop_app:
	docker-compose down go-url-shortner-application