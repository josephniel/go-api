# go-api

## Description
`go-api` is a boilerplate for RESTful microservices written in Go. This is a wrapper for the [echo famework](https://github.com/labstack/echo) which aims to simplify and further structurize API development in Go.

## Notice
This repository is still a work in progress.

## Installation
The application is dockerized. To be able to run, you can:
```
docker-compose build
docker-compose up
```
to build and spin up the dockerized application.

## Directory Structure
- `/app`
    - `/bootstrap` contains all the application preparations (i.e. getting environment configuration, setting up middlewares, and registering routers)

    - `/config` contains the base configuration structure used by the application when instantiating. This also contains environment-specific configuration named: `<env>.yaml`. 

    - `/controller` contains all the entrypoints in the application. A common practice that should be followed in this repository is for every controller, we only need to call the helper functions `Get`, `Post`, `Patch` in the `router` package to register the route for the particular method selected. We can isolate the route definition for the specified controller by defining an `init()` function in the file.

    ```
    package controller

    import "github.com/josephniel/go-api/app/router"

    func sampleController(context *echo.Context) error {
        ...
    }

    ...

    func init() {
        router.Get(<relative path>, sampleController)
        ...
    }
    ```
    The relative path follows the path specification in Echo.

    - `/model` contains database mapping for the application. For this repository, we use postgres with the [pg ORM](https://github.com/go-pg/pg). The defined structure also doubles as a schema for requests such as `POST` or `PATCH`. This is to reduce redundancy when defining structures for a model and a schema. Also, the database connection is accessible in `model.DB`. All database operations should be done in the `/operation` folder using the `model.DB` connection.

    - `/operation` contains all the business logic performed in the controller. This is directly called by the controller and should always be used. Business logic should never be in the controller (i.e. as it gets cluttered really easy)

    - `/router` contains the helper methods for registering a route for a controller.

    - `/schema` contains structures that work as schemas/DTOs for the application.

    - `/util` contains general utility functions.

    - `main.go` is the entrypoint of the application.
