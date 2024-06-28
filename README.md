docker compose up -d

docker compose ps

docker exec -it appproduct bash

go mod init github.com/uiratan/fullcycle-archdev-hexagonal

go test ./...

mockgen -destination=application/mocks/application.go -source=application/product.go application