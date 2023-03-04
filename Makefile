
run-docker:
	docker-compose up -d --force-recreate --quiet-pull 
run-local:
	go run main.go
kill:
	kill -9 (lsof -i :8080 | grep PID)
post: 
	curl -X POST -H "Content-Type: application/json" \
    -d '{"username": "thisisatest", "email": "ajalantbrown@yahoo.com", "password": "password1234", "dob":"01/01/2021"}' \
    http://localhost:8080/register

login: 
	curl -X POST -H "Content-Type: application/json" \
    -d '{"username": "thisisatest", "password": "password1234"}' \
    http://localhost:8080/login
up:
	migrate -path packages/database/migrations -database "postgresql://postgres:postgres@localhost:5432/libre?sslmode=disable" -verbose up
down:
	migrate -path packages/database/migrations -database "postgresql://postgres:postgres@localhost:5432/libre?sslmode=disable" -verbose down