BROKER_BINARY=brokerApp
FRONT_END_BINARY=frontApp
AUTH_BINARY=authApp
LOGGER_BINARY=loggerApp

start-containers:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

stop-containers:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"


.PHONY: clean-docker-images
clean-docker-images:
	@echo "Removing all Docker images..."
	docker rmi $(shell docker images -q)
	@echo "All Docker images removed!"

build-broker:
	@echo "Building broker binary..."
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"

## build-auth: builds the auth binary as a linux executable
build-auth:
	@echo "Building auth binary..."
	cd ../authentication-service && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd/api
	@echo "Done!"

## build-logger: builds the logger binary as a Linux executable
build-logger:
	@echo "Building logger binary..."
	cd ../logger-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LOGGER_BINARY} ./cmd/api
	@echo "Done!"

## build-front: builds the front end binary
build-front:
	@echo "Building front end binary..."
	cd ../front-end && env CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} ./cmd/web
	@echo "Done!"

## start-front: starts the front end
start-front: build-front
	@echo "Starting front end"
	cd ../front-end && ./${FRONT_END_BINARY} > /dev/null 2>&1 &

## stop-front: stop the front end
stop-front:
	@echo "Stopping front end..."
	@-pkill -SIGTERM -f "./${FRONT_END_BINARY}"
	@echo "Stopped front end!"

rebuild-containers: build-broker build-auth build-logger
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"