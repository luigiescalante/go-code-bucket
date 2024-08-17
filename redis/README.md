# GO Redis Crud
An example for use redis to save, get, update and delete data.

![GitHub tag (with filter)](https://img.shields.io/github/v/tag/luigiescalante/go-api-template)
![GitHub last commit (branch)](https://img.shields.io/github/last-commit/luigiescalante/go-api-template/main)
![Static Badge](https://img.shields.io/badge/email-luigi.escalante%5Bat%5Dgmail.com-blue)
![X (formerly Twitter) Follow](https://img.shields.io/twitter/follow/luigi_escalante)
<p align="center">
<img src="../github-logo.png" alt="logo" width="200" height="292">
</p>

# Redis Code examples
- Save a new entity
- Update an entity
- Get one entity
- Delete an entity

## Requirements
- Go
- Docker
- Docker Compose
## Install project
1. Rename the file .env.example to .env and fill the paramters
2. Start docker compose for database redis
~~~~
    sudo docker-compose up -d --build
~~~~
3.Set env vars on your environment
~~~~
export REDIS_ADDRESS=<host (Ex. localhost)>
export REDIS_PORT=<user database (Ex. 6379)>
export REDIS_PASSWORD=<redis password (Ex. 123456)>
export REDIS_DB=<redis database (Ex. 0)>
~~~~

## Information and tutorial

https://dev.to/luigiescalante/go-bun-orm-with-postgresql-quickly-example-394o