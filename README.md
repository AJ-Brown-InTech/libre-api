# Libre (Social Media) Rest Api

### golang v1.18

# Reference Makefile for commands
## 

Start the containers: `docker-compose up -d`

Rebuild and start: `docker-compose up -d --build`

Login: `docker exec -it postgres psql -U user`

This command will be determined by the docker-compose.yml file if different env variables are used.


find dcoker ip```ip addr show docker0 | grep -Po 'inet \K[\d.]+'```