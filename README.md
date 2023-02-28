# Libra (Social Media) Rest Api

### This is a rest api for *Libra* a social media mobile application for IOS.

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
    - *utils*
    - *middleware*
    - *models*
    - *routes*
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
        - You can also pull down the docker container which is also public in the registry.
    
## If you would like to get involved please feel free reach out via email at ``` ajalantbrown@gmail.com```