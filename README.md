# Hotel Reservation CRUD API

## Project

- users -> book room from an hotel
- admins -> going to check reservation/bookings
- Authentication and authorization -> JWT tokens
- Hotels -> CRUD API -> JSON
- Rooms -> CRUD API -> JSON
- Scripts -> datasabe management -> seeding, migration

## Resources

### MongoDb

Documenation

```url
https://www.mongodb.com/docs/drivers/go
```

Install

```bash
go get go.mongodb.org/mongo-driver/mongo
```

### gofiber

Documentation

```url
https://docs.gofiber.io
```

Install

```bash
go get github.com/gofiber/fiber/v2
```

### Docker

## Install mongodb as Docker container

```bash
docker run --name mongodb -d mongo:latest -p 27017:27017
```

## Run the API as Docker container

````bash
make docker
````
