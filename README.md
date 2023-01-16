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

---
- **Packages**
    - *database*
        1. database configurations
            : Package for connecting to postgres database.
    - *middleware*
        1. middlware
            : Creates a cookie sesssion.
            : Verifies a cookie sesssion.
    - *utils*
        1. logs
            : custom logging for REST api.
---
- **Enviroment**
    - *Development*
        : Currently only utilizes a develpoment enviroment.
        : Development Configuration file is located "**/config/config-local.yml**"
    - *Running Api*
        : Simply execute ``` make libre ``` to run the api
      
 
