ngrok:
	ngrok http http://localhost:8080

postgres:
	docker run --name heatmap_db -p 5432:5432 \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=secret \
		-v pgdata:/var/lib/postgresql/data \
		-d postgres:17.2-alpine

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
	docker rm -f heatmap_api heatmap_db || true
	docker compose up --builddocke

run:
	docker container rm heatmap-api && docker run --name heatmap-api -d -p 8080:8080 heatmap-api 

.PHONY: ngrok postgres postgres-down createdb migrateup sqlc build composeup