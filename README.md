# Event Sourcing & CQRS Pattern

Introduction to Event sourcing and CQRS application architecture patterns into practice.

## Stack
- go
- mysql
- redis
- docker

## Project setup (Docker)
1- Clone the repository

2- Run `docker-compose up --build` to build the containers.

3- Import the `CQRS-ES.postman_collection.json` file in postman in order to test the routes.

> You can use [redis insight](https://redislabs.com/redis-enterprise/redis-insight/) to check the events stored in redis.  

> Conn parameter: Name: cqrs, Port: 6379, Password: cqrs

> For now the only event stored is the order created event.

Ports: 

- MySQL: `3306`
- Redis: `6379`
- GoApp: `8000`

## Note
CQRS is implemented with the user's commands and queries.

Event Sourcing is not yet working, the only working side is that the order created event is stored in redis streams and a snapshot is created. 