
run:
	docker-compose up -d --force-recreate --quiet-pull 
#manual run
man:
	go run main.go
kill:

	kill -9 (lsof -i :8080 | grep PID)
post: 
	curl -X POST -H "Content-Type: application/json" \
    -d '{"username": "thisisatest", "email": "ajalantbrown@yahoo.com", "password": "password1234", "birthdate":"01/01/2021"}' \
    http://localhost:8080/register