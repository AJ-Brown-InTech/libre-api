
run:
	docker-compose up -d --force-recreate --quiet-pull 
#manual run
man:
	go run main.go
kill:

	kill -9 (lsof -i :8080 | grep PID)
post: 
	curl -X POST -H "Content-Type: application/json" \
    -d '{"username": "123456", "email": "email@gmail.com", "password": "password", "birthdate":"01/01/2021", "firstName": "first", "lastName":"last"}' \
    http://localhost:8080/register