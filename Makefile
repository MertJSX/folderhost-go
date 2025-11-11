start:
	go run main.go
web:
	cd web && npm run dev
# Use docker for development if you want...
docker:
	docker compose build
	docker compose up
# Use mingw for to take Windows build. The other versions won't work :/
build: 
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o ./debug/folderhost main.go
	CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o ./debug/folderhost.exe main.go
setup:
	@echo "Downloading dependencies..."
	go mod tidy
	go mod download
	cd web && npm install

	@echo "Dependencies are downloaded successfully."