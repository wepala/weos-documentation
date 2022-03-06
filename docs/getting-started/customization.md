# WeOS Custom Binary

If you know a bit about Go, you can use WeOS as a framework for creating REST API using OpenAPI. In WeOS middleware, controllers and initializers are simply functions you can create and store in the dependency container. Dependencies are registered can be referenced in the OpenAPI specification. Projections, event stores, and command dispatchers can all be customized.

## Starting A Customized  API

To get started with building your own custom Go application using WeOS

1. Install `go get github.com/wepala/weos`
2. Setup api
```go
package main

import (
	weos "github.com/wepala/weos/controllers/rest"
	"golang.org/x/net/context"
	"log"
)

func main() {
	ctxt := context.Background()
	//instantiate weos with a reference to the OpenAPI specification
	api, err := weos.New("api.yaml")
	if err != nil {
		log.Fatalln("error loading api config", err)
	}
   //initialize API so that standard middleware,controller,projections etc are registered
	err = api.Initialize(ctxt)
	if err != nil {
		log.Fatalln("error initializing api", err)
	}
	//start API 
	weos.Serve("8681", api)
}

```

## Adding Custom Middleware

Any middleweare must satisfy the signature of `rest.Middleware` e.g.

```go 
func CustomMiddleware(api *RESTAPI, projection projections.Projection, commandDispatcher model.CommandDispatcher, eventSource model.EventRepository, entityFactory model.EntityFactory, path *openapi3.PathItem, operation *openapi3.Operation) echo.MiddlewareFunc {
	//TODO code you want to execute once can go here e.g. parsing information from OpenAPI spec
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctxt echo.Context) error {
			//TODO code that you want to execute per request goes here. 
			//variables define in the OpenAPI spec will be available in the request context ctxt.Request().Context()
		}
	}
}
```

To see the full list of information automatically available in the context see the documentation of the context

Once the middleware is defined it must be registered with the dependency container

```go
package main

import (
	controllers "github.com/wepala/weos/controllers/rest"
	projections "github.com/wepala/weos/projections"
	model "github.com/wepala/weos/model"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"
	"log"
)

func CustomMiddleware(api *controllers.RESTAPI, projection projections.Projection, commandDispatcher model.CommandDispatcher, eventSource model.EventRepository, entityFactory model.EntityFactory, path *openapi3.PathItem, operation *openapi3.Operation) echo.MiddlewareFunc {
	//TODO code you want to execute once can go here e.g. parsing information from OpenAPI spec
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctxt echo.Context) error {
			//TODO code that you want to execute per request goes here. 
			//variables define in the OpenAPI spec will be available in the request context ctxt.Request().Context()
			return nil
		}
	}
}

func main() {
	ctxt := context.Background()
	//instantiate weos with a reference to the OpenAPI specification
	api, err := controllers.New("api.yaml")
	if err != nil {
		log.Fatalln("error loading api config", err)
	}
   //register middleware
	api.RegisterMiddleware("Custom",CustomMiddleware)
   //initialize API so that standard middleware,controller,projections etc are registered
	err = api.Initialize(ctxt)
	if err != nil {
		log.Fatalln("error initializing api", err)
	}
	//start API 
	controllers.Serve("8681", api)
}

```

With the middleware registered you can now reference it in the OpenAPI spec

```yaml
openapi: 3.0.3
info:
  title: Todo Service
  description: Some Todo Service
  version: 1.0.0
servers:
  - url: 'http://localhost:8681'
x-weos-config:
  database:
    driver: sqlite3
    database: weprojects.db
components:
  schemas:
    Task:
      type: object
      properties:
        title:
          type: string
        description:
          type: string
        type:
          type: string
          enum:
            - Story
            - Task
        dueDate:
          type: string
          format: date-time
paths:
  /tasks:
    get:
      operationId: getTasks
      x-middleware:
        - Custom
      responses:
        200:
          description: List of tasks
          content:
            application/json:
              schema:
                type: object
                properties:
                  total:
                    type: integer
                  page:
                    type: integer
                  items:
                    type: array
                    items:
                      $ref: "#/components/schemas/Task"
    post:
      operationId: createTask
      requestBody:
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/Task"
      responses:
        201:
          description: Created task
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Task"
```