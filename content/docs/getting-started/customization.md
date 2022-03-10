---
title: Customization
weight: 20
description: Build a custom WeOS binary.
---

If you know a bit about Go, you can use WeOS as a library for creating REST API using OpenAPI (as opposed to using the weos-cli). In WeOS middleware,
controllers and initializers are simply functions you can create and store in the dependency container. Dependencies are
registered can be referenced in the OpenAPI specification. Projections, event stores, and command dispatchers can all be
customized.

## Starting A Customized API

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
3. Compile the app using `go build`
4. Run api by calling the app that was compiled in the previous step

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

Once the middleware is defined it must be registered with the dependency container.

```go
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

With the middleware registered you can now reference it in the OpenAPI spec e.g.
```yaml
/tasks:
  get:
    operationId: getTasks
    x-middleware:
      - Custom
```
[Complete Example](../examples/customizations/custom_middleware)


## Customizing Projections

WeOS provides a default projection "Default" that supports a few databases (sqlite, Postgresql, Mysql, MariaDB). There are
a few reasons for creating custom projections including:

1. Using an unsupported database
2. Using events to generate data that should be re-created whenever the events are played (e.g. using events to generate QRcode images)

In order to create your own custom projection you must create a struct that implements `weos.Projection`

```go
//CustomProjection example of a custom projection
type CustomProjection struct {
}

func (c *CustomProjection) Migrate(ctx context.Context, builders map[string]dynamicstruct.Builder, deletedFields map[string][]string) error {
    return nil
}

func (c *CustomProjection) GetEventHandler() model.EventHandler {
    return func(ctx context.Context, event model.Event) error {
        switch event.Type {
            default:
                println("event handler hit")
        }
        return nil
    }
}

func (c *CustomProjection) GetContentEntity(ctx context.Context, entityFactory model.EntityFactory, weosID string) (*model.ContentEntity, error) {
//TODO implement me
panic("implement me")
}

func (c *CustomProjection) GetByKey(ctxt context.Context, entityFactory model.EntityFactory, identifiers map[string]interface{}) (map[string]interface{}, error) {
    entity := make(map[string]interface{})
    entity["title"] = "Foo"
    entity["description"] = "Bar"
    return entity, nil
}

func (c *CustomProjection) GetByEntityID(ctxt context.Context, entityFactory model.EntityFactory, id string) (map[string]interface{}, error) {
    entity := make(map[string]interface{})
    entity["title"] = "Foo"
    entity["description"] = "Bar"
    return entity, nil
}

func (c *CustomProjection) GetContentEntities(ctx context.Context, entityFactory model.EntityFactory, page int, limit int, query string, sortOptions map[string]string, filterOptions map[string]interface{}) ([]map[string]interface{}, int64, error) {
    entity := make(map[string]interface{})
    entity["title"] = "Foo"
    entity["description"] = "Bar"
    return []map[string]interface{}{entity}, 1, nil
}

func (c *CustomProjection) GetByProperties(ctxt context.Context, entityFactory model.EntityFactory, identifiers map[string]interface{}) ([]map[string]interface{}, error) {
    //TODO implement me
    panic("implement me")
}
```
!!! Tip
The event handler typically has a switch that handles different event types

To make the projection accessible via the open api spec it needs to be registered on the dependency container

```go
func main() {
    ctxt := context.Background()
    //instantiate weos with a reference to the OpenAPI specification
    api, err := controllers.New("api.yaml")
    if err != nil {
        log.Fatalln("error loading api config", err)
    }
    //register projection
    api.RegisterProjection("Default",CustomProjection)
    //initialize API so that standard middleware,controller,projections etc are registered
    err = api.Initialize(ctxt)
    if err != nil {
        log.Fatalln("error initializing api", err)
    }
    //start API 
    controllers.Serve("8681", api)
}

```

!!! Tip
Projection event handlers should be done in an idempotent way to accommodate replaying of events during a system recovery

### Default projection
The default projection can be replaced by registering the custom projection as the `Default` projection in the container
```go
//register middleware
api.RegisterProjection("Default",CustomProjection)
```

### Adding Multiple Projections
Multiple projections can be associated with an operation by using the `x-projections` extension.
```yaml
/tasks:
  get:
    operationId: getTasks
    x-projections:
      - Custom
      - Default
```

The order of the projections matters, especially for endpoints that return data from the projection. In the case there are
multiple projections, WeOS will loop through each projection until it gets non nil result. If there really was no result
WeOS would have made a call to each configured projection.

[Complete Example](../examples/customizations/multiple_projections)
