### Golang My Bank Api 🚀

#### 👨‍💻 Full list what has been used:
* [fiber](https://gofiber.io/) - Web framework
* [sqlx](https://github.com/jmoiron/sqlx) - Extensions to database/sql.
* [pgx](https://github.com/jackc/pgx) - PostgreSQL driver and toolkit for Go
* [uuid](https://github.com/google/uuid) - UUID
* [migrate](https://github.com/golang-migrate/migrate) - Database migrations. CLI and Golang library.
* [swag](https://github.com/swaggo/swag) - Swagger
* [testify](https://github.com/stretchr/testify) - Testing toolkit
* [gomock](https://github.com/golang/mock) - Mocking framework
* [CompileDaemon](https://github.com/githubnemo/CompileDaemon) - Compile daemon for Go
* [Docker](https://www.docker.com/) - Docker

#### Migration:
    make migrate_up
    make migrate_down

#### Recomendation for local development most comfortable usage:
    make develop // run all containers
    make run

#### 🙌👨‍💻🚀 Docker-compose files:
    compose.yml - run docker development environment

### Docker development usage:
    make develop
    make down
    make stop
    make test
    make doc-generate

### Local development usage:
    make build
    make run

### SWAGGER UI:

https://localhost:8080/swagger/index.html
