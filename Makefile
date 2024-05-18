migrateup:
	docker run --rm -v $(shell pwd)/db/migration:/migration --network host migrate/migrate -path=/migration/ -database "postgresql://postgres:postgres@localhost:5432/se_school?sslmode=disable" -verbose up

migratedown:
	docker run --rm -v $(shell pwd)/db/migration:/migration --network host migrate/migrate -path=/migration/ -database "postgresql://postgres:postgres@localhost:5432/se_school?sslmode=disable" down -all

sqlc:
	docker run --rm -v $(PWD):/app -w /app kjconroy/sqlc generate

test:
	go test -v -cover ./db/sqlc

server:
	go run cmd/main.go

swag:
	${HOME}/go/bin/swag init -d cmd,api,db/sqlc

.PHONY: migrateup migratedown sqlc test server swag