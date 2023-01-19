run-api:
	docker pull ajalanbrown/libre-api
	docker run --platform linux/amd64 -it ajalanbrown/libre-api bash --net host
cont:
	docker pull ajalanbrown/libre-api
	cd docker/
	docker compose up 
	docker-compose exec web sh 
run:
	docker-compose up -d --force-recreate --quiet-pull 
#manual run
man:
	go run main.go
kill:

	kill -9 (lsof -i :8080 | grep PID)

	