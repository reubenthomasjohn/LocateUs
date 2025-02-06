ngrok:
	ngrok http http://localhost:8080

postgres:
	docker run --name heatmap_db -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -d postgres:17.2-alpine

postgres-down:
	docker stop heatmap_db && docker container rm heatmap_db

createdb:
	docker exec -it heatmap_db createdb --username=postgres --owner=postgres ccc

migrateup:
	migrate -path db/migration -database "postgres://postgres:secret@localhost:5432/ccc?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://postgres:secret@localhost:5432/ccc?sslmode=disable" -verbose down

sqlc:
	sqlc generate

build:
	docker build --tag heatmap-api .

composeup:
	docker compose up

.PHONY: ngrok postgres postgres-down createdb migrateup sqlc build composeup