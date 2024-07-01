docker compose up -d
docker compose ps
docker exec -it appproduct bash

go mod init github.com/uiratan/fullcycle-archdev-hexagonal

go test ./...

mockgen -destination=application/mocks/application.go -source=application/product.go application

touch sqlite.db
sqlite3 sqlite.db
create table products(id string, name string, price float, status string);
.tables

cobra-cli

go run main.go http
