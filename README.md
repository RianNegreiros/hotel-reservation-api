# Lodging App

## Application inspired on Airbnb to search for short-term rental

## [Live Site](https://go-lodging.onrender.com/)

### Preview

![Application Preview 1](/Documentation/Preview1.png)
![Application Preview 2](/Documentation/Preview2.png)

### Methodologies

- MVC

### Tecnologies and Tools

- [Golang](https://go.dev/)
- [PostgresQL](https://www.postgresql.org/)
- [MailHog](https://github.com/mailhog/MailHog)
- [nosurf](https://pkg.go.dev/github.com/justinas/nosurf@v1.1.1)
- [SCS](https://pkg.go.dev/github.com/alexedwards/scs/v2@v2.5.0)

## How to run

### Prerequisites

[Docker Compose](https://docs.docker.com/compose/gettingstarted/)
[Soda](https://gobuffalo.io/documentation/database/soda/)

1. Remove **.example** from run.exmaple.sh filename, fill the parameters:
  - dbhost(default:localhost), dbname, dbuser, dbpass, dbport(default: 5432), dbssl(default: disable) according to PostgreSQL config on docker-compose.yml
  - useCache(default:true) to cache the application templates
  - inProduction(default:true) to set prodcution mode
3. Remove **.example** from database.example.yml filename and fill the values according to PostgreSQL config on docker-compose.yml

4. Run soda migrate

5. Execute ./run.sh

6. Go to localhost:8080 on your preferably browser
