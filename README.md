# Libre (Social Media) Rest Api

### This is a rest api for *libre* a social media mobile application for IOS.

| Section | Description |
| ----------- | ----------- |
| Tools | Versions/Frameworks/ |
| Packages | Api Architecture |
| Enviroment | Local Development Enviroment

---
- **Tools**
    1. [Golang(version 1.18)](https://go.dev/)
    2. [Fiber(V2)](https://docs.gofiber.io/)
    3. [Docker](https://docs.docker.com/get-docker/)
    4. [Postgres](https://www.postgresql.org/download/)
    5. [Air(Server hot reload)](https://github.com/cosmtrek/air)
    6. [Migrate](https://github.com/golang-migrate/migrate)

---
- **Packages**
    - *database*
        1. database configurations
            : 
            - Package for connecting to postgres database.
        1. migrations 
            : 
            - database migrations
            -  ex.(migrate create -ext sql -dir database/migrations/ -seq init_create_users)
    - *middleware*
        1. middlware file
            : 
              - Creates a cookie sesssion.
              - Verifies a cookie sesssion.
    - *utils*
        1. logs
            : 
             - custom logging for REST api.
---
- **Enviroment**
    - *Development*
        : 
        - Currently only utilizes a develpoment enviroment.
        - Development Configuration file is located "**/config/config-local.yml**"
    - *Running Api*
        : 
        - Simply execute ``` make libre ``` to run the api manually
        - or run ```air``` 
    - *Weird*
        :
        - running air and not closing the server ``` CTRL C``` when closing leaves server running and will have to manually run ``` lsof -i :8080 ``` then whatever the PID is pass it into ```kill -9 <PID>``` this kills the proccess if using **air**

## If you would like to get involved please reach out via email at ``` ajalantbrown@gmail.com```
 
