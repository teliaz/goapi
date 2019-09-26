# GoLang API

## Introduction

Since I'm new to GO programming language I took some time to scratch on the basics

- [Go by Example](https://gobyexample.com)
- [A tour of Go](https://tour.golang.org)

Then dove in on the following

- [Functional Programming in Go](https://medium.com/@geisonfgfg/functional-go-bc116f4c96a4)
- [Solid Design in GO](https://dave.cheney.net/2016/08/20/solid-go-design)
- [SQLite Concurrency Issue-fix](https://itnext.io/telegram-bot-in-go-concurrent-sqlite-e6176fac088e) trying to avoid using multiple containers with a real storage technology.
- [Interfaces and composition for effective testing](https://nathanleclaire.com/blog/2015/10/10/interfaces-and-composition-for-effective-unit-testing-in-golang)
- [7 Mistakes to avoid in Go](https://www.youtube.com/watch?v=29LLRKIL_TI)
- [Never Use (Mat Ryer)](https://www.youtube.com/watch?v=5DVV36uqQ4E)
- [Comparing go web frameworks](https://github.com/diyan/go-web-framework-comparsion)
- [Go frameworks pros and cons](https://nordicapis.com/7-frameworks-to-build-a-rest-api-in-go/)
- [Dockerized Implementation with Postgress](https://github.com/kisulken/bulletinApi/blob/master/main.go)
- [How to automatically handle vendor folder with godep](https://github.com/tools/godep)
- [How to structure your Go Applications](https://www.youtube.com/watch?v=VQym87o91f8)
- [Concurrency made easy](https://www.youtube.com/watch?v=DqHb5KBe7qI)
- [GopherCon2017 - Understanding Channels](https://www.youtube.com/watch?v=KBZlN0izeiY)

## Docker Guideline

To run the API use the following:

```bash
docker-compose -f ./docker-compose.yml up -d
docker-compose build up
```

Dockerfile included in the project

## API Routes table

Endpoint                                        | Description
------------                                    | -------------
(GET)/assets                                    | endpoint to receive a user id and return a list of all the user’s assets
(POST)/assets/{id}/favorite/{bool}              | endpoints that would add an asset to favourites, remove it
(PUT)/assets/{id}/description                   | edit its description

Challenge Keypoints

- [x] A working server application with functional API is required
- [x] It is appreciated, though not required, if a Dockerfile is included.
- [ ] Note that users have no limit on how many assets they want on their favourites so your service will need to provide a reasonable response time
- [ ] Useful and passing tests would be also be viewed favourably

Features

- [x] JWT Authentication, Authorizations
- [ ] ORM Implementation
- [ ] Testable Endpoints

External Packages Used

- [x] Gorilla Mux - Router
- [x] GORM - A Go ORM
- [x] Dgrijalva JWT - JSON Web Token Library

Some other resources

- [Docker Compose Settings](https://github.com/kisulken/bulletinApi/blob/master/docker-compose.yml)
- [Initial Implementation](https://github.com/dedidot/simple-api-golang)
- [Decorators example](https://gist.github.com/thomasdarimont/31b26f782644c92effd0df3f7b64ef5d)
- [Channels Implementation](https://www.youtube.com/watch?v=7DXQH7bMvZ8)
- [Simple-API-Golang](github.com/mingrammer/go-todo-rest-api-example)
- [CRUD RESTful API with Go, GORM, JWT, Postgres, Mysql, and Testing](https://github.com/victorsteven/Go-JWT-Postgres-Mysql-Restful-API)
