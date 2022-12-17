run-dock:
	docker pull ajalanbrown/libre-api
	docker run --platform linux/amd64 -it ajalanbrown/libre-api bash
run-both:
	docker pull ajalanbrown/libre-api
	cd docker/
	docker compose up 
	docker-compose exec web sh 