# go-api

## Description
`go-api` is a boilerplate or RESTful microservices written in Go. This is a wrapper for the [echo famework](https://github.com/labstack/echo) which aims to simplify and further structurize API development in Go.

## Notice
This repository is still a work in progress.

## Installation
This is meant to be cloned and not imported as a package; You can `go get github.com/josephniel/go-api` and just replace all instances for the imports relative to the repository namespace.

After setting up, you can try to load the application by running
```
go run app/main.go
```
You can also set an environment flag in the command (i.e. to be able to read environment-specific configurations for the application).
```
go run app/main.go -env=<env>
```

## Directory Structure
- `/app`
    - `/bootstrap` contains all the application preparations (i.e. getting environment configuration, setting up middlewares, and registering routers)

    - `/config` contains the base configuration structure used by the application when instantiating. This also contains environment-specific configuration named: `<env>.yaml`. 

    - `/controllers` contains all the entrypoints in the application. A common practice that should be followed in this repository is for every controller, we only need to call the helper functions `Get`, `Post`, `Patch` in the `router` package to register the route for the particular method selected. We can isolate the route definition for the specified controller by defining an `init()` function in the file.

    ```
    package controllers

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

    - `/models` should *ideally* only contain structures for the database mapping. 

    - `/operations` contains all the business logic performed in the controller. This is directly called by the controller and should always be used. Business logic should never be in the controller (i.e. as it gets cluttered really easy)

    - `/router` contains the helper methods for registering a route for a controller.

    - `/schema` contains structures that work as schemas/DTOs for the application.

    - `/utils` contains general utility functions.

    - `main.go` is the entrypoint of the application.
