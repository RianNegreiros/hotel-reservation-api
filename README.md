# Hotel Reservation Back-end

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

```
https://www.mongodb.com/docs/
```

Install

```bash
go get go.mongodb.org/mongo-driver/mongo
```

### gofiber

Documentation

```
https://gofiber.io
```

Install

```bash
go get github.com/gofiber/fiber/v2
```

## Docker

### Install mongodb as Docker container

```bash
docker run --name mongodb -d mongo:latest -p 27017:27017
```
