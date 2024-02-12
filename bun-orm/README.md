# GO Bun ORM with Postgresql
An examples collection for use Bun, a light ORM for postgresql database.

![GitHub tag (with filter)](https://img.shields.io/github/v/tag/luigiescalante/go-api-template)
![GitHub last commit (branch)](https://img.shields.io/github/last-commit/luigiescalante/go-api-template/main)
![Static Badge](https://img.shields.io/badge/email-luigi.escalante%5Bat%5Dgmail.com-blue)
![X (formerly Twitter) Follow](https://img.shields.io/twitter/follow/luigi_escalante)
<p align="center">
<img src="../github-logo.png" alt="logo" width="200" height="292">
</p>

# Bun Code examples
- Save a new entity
- Update an entity
- Soft delete entity
- Get one entity
- Get an entity list

## Requirements
- Go
- Docker
- Docker Compose
## Install project
1.  Start docker compose for database postgresql
~~~~
    sudo docker-compose up -d --build
~~~~
2.  Set env vars on your environment
~~~~
export DB_USER=<user database (Ex. admin)>
export DB_PASSWORD=<user password (Ex. admin123)>
export DB_NAME=<db name (Ex. code_bucket)>
export DB_HOST=<host address (Ex. localhost)>
export DB_PORT=<database port (Ex. 5432)>
~~~~

## Information and tutorial

https://dev.to/luigiescalante/go-bun-orm-with-postgresql-quickly-example-394o