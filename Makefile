DB_URL := "postgres://postgres:+_+diyor2005+_+@localhost:5432/book?sslmode=disable"
run:
	go run cmd/soldiers_service/main.go

migrate-file:
	migrate create -ext sql -dir migrations/ -seq book_service

migrate-up:
	migrate -path migrations -database $(DB_URL) -verbose up

migrate-down:
	migrate -path migrations -database $(DB_URL) -verbose down


migrate-force:
	migrate -path migrations -database $(DB_URL) -verbose force 1

SWAGGER_CMD = swag

swagger-init:
	$(SWAGGER_CMD) init -g api/handler/handler.go -o api/docs
