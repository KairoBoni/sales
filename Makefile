BUILD_FOLDER=bin

build_backend:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o $(BUILD_FOLDER)/api ./cmd/api

run_backend:
	CONFIG_FILEPATH=./backend/config.yaml go run ./backend/.