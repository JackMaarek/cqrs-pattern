#Event Sourcing & CQRS Pattern

Introduction to Event sourcing and CQRS application architecture patterns into practice.

##Stack
- go
- mysql
- redis
- docker

##Project setup
1- Clone the repository

2- Run `docker-compose up --build` to build the containers

3- Enter the app container with `docker-compose exec app /bin/sh` and then compile the go app with `go build -o main .`

4- Run the application with the binary `./main`

> You need to build the project each time a change has been done.