# final

You need go (golang) version 1.14

1. chmod +x migrate
2. chmod +x migrate.sh
3. docker-compose up -d --build (it's run database)
4. ./migrate.sh
5. go mod download
6. go build -o app ./cmd/main.go
7. ./app

server (API) available at localhost:8000
