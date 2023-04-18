run:
	go run ./cmd/api

dup:
	docker-compose up --build -d

down:
	docker-compose down