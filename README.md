# GlobalWebIndex Engineering Challenge - Elias Krontiris Submission

## Problem Analysis & Concerns

After reading carefully the challenge goals all assets (charts, insights, audience) could generate data from a single sample data-table. Guessing the Age, Gender, Country are `demographic information` and the only `metric` is the "hours spent on social media", I added the "Participants" schema and seeded mock data into this table.
Age(10-80), Gender(2) and Country(259) are randomly seeded values and "hours spent on social media" is based on a normal distribution with a spike on 0-2 hours daily, based on a personal estimation.

I should note that I fully understand that this is out of this challenge's scope and lacks performance, since all *Assets* are calculated based on live data and usually should be exported based on pre-calculated and transformed schemas.
Also noting that using a ORM deplets the neccesity to use repository pattern for storing data.

For convinience and easy testing the endpoint on every spin-up the DB is dropped and recreated with mock data.
I also include a collection of API calls test all chalenges endpoints [Postman Collection Link](https://www.getpostman.com/collections/dd8e929f0dd1124fbb3a)

## Docker Spinup Guidelines

To run the API use the following:

```bash
docker-compose -f ./docker-compose.yml up
docker-compose up -d --force-recreate --build
```

Dockerfile included in the project


## API Routes table

Endpoint                                        | Description
------------                                    | -------------
(POST)/auth/signup                              | endpoint to create a user
(POST)/auth/login                               | authorize and return json web token
(GET)/assets                                    | based on the token returns all user's assets
(GET)/assets/{id}                               | get specific asset by id
(PATCH)/assets/{id}                             | edit asset's description or isFavorite
(POST)/assets/charts                            | create asset / chart
(POST)/assets/insigts                           | create asset / insight
(POST)/assets/audiences                         | create asset / audience



there are more endpoints that were made to make Seeding and Testing easier. All routes are specified on `/app/routes.go`

## Challenge Keypoints

- [x] A working server application with functional API is required
- [x] It is appreciated, though not required, if a Dockerfile is included.
- [ ] Note that users have no limit on how many assets they want on their favourites so your service will need to provide a reasonable response time
- [ ] Useful and passing tests would be also be viewed favourably *(Incomplete Tests)

## Features

- [x] JWT Authentication, Authorizations
- [x] ORM Implementation
- [x] Mock Data
- [ ] Testing Implementations *(Incomplete Tests)

## Dependencies

- [x] Gorilla Mux - Router
- [x] GORM - A Go ORM
- [x] Dgrijalva JWT - JSON Web Token Library

## Some other resources

- [Docker Compose Settings](https://github.com/kisulken/bulletinApi/blob/master/docker-compose.yml)
- [Initial Implementation](https://github.com/dedidot/simple-api-golang)
- [Decorators example](https://gist.github.com/thomasdarimont/31b26f782644c92effd0df3f7b64ef5d)
- [Channels Implementation](https://www.youtube.com/watch?v=7DXQH7bMvZ8)
- [Simple-API-Golang](github.com/mingrammer/go-todo-rest-api-example)
- [CRUD RESTful API with Go, GORM, JWT, Postgres, Mysql, and Testing](https://github.com/victorsteven/Go-JWT-Postgres-Mysql-Restful-API)
