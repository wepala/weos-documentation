---
title: Middleware
description: Reusable code for handling cross-cutting concerns.
---

Middleware is reusable code that can be associated with a path or set globally.

## Global Middleware

Use the `x-weos-config` extension to configure global middleware:
```yaml
x-weos-config:
  rest:
    middleware:
      - RequestID
      - Recover
      - Logger
      - ZapLogger
      - CSVUpload
```
### Standard Global Middleware 

- **RequestID**: generate request IDs
- **Recover**: recover from panics and handle appropriately
- **Logger**: log information about HTTP requests
- **ZapLogger**: log using structured, leveled logging
- **CSVUpload**: parse CSV files

## Path middleware

Path middleware contain logic that you may want to execute per request. To configure middleware on a path:
```yaml
paths:
  /tasks:
      get:
        operationId: createTasks
        x-middleware:
          - CreateMiddleware
```

### Standard Path Middleware

To make it easier to do Create, Read, Update, Delete (CRUD) functionality, there are standard middlewares that you can use
on a path. Depending on how you setup your API specification WeOS will automatically associate middleware with your paths
if you don't specify a controller. As much as possible you should try to use standard middleware (why write more code than
you have to?)

| Middleware            | Description                                                                                                                       | Context Constant(s) | Context Value Key  |
|:----------------------|:----------------------------------------------------------------------------------------------------------------------------------|:-------------------:|:-------------------|
 | CreateMiddleware      | Creates an entity using the path schema and request payload                                                                       |      ENTITY_ID      | _entity_id         |
 | CreateBatchMiddleware | Creates entities using the path schema and request payload                                                                        |                     |                    |
 | ViewMiddleware        | View entity by specifying user defined id OR WeOS defined entity id. Can also be used to retreive a specific version of an entity |       ENTITY        | _entity            |
 | ListMiddleware        | View a list of entities                                                                                                           |  ENTITY_COLLECTION  | _entity_collection |
 | UpdateMiddleware      | Update entity                                                                                                                     |      ENTITY_ID      | _entity_id         |
 | DeleteMiddleware      | Delete entity                                                                                                                     |      ENTITY_ID      | _entity_id         |


## Adding Custom Middleware

Any middleware must satisfy the signature of `rest.Middleware` e.g.

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
