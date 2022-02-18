---
title: Setting up routes
---

API routes are what applications use to access data and execute functionality. The paths you specify in the [OpenAPI
specification](https://swagger.io/docs/specification/paths-and-operations/) will become endpoints to which your
application can send requests. Each path can have multiple operations that you can configure separately or you can
configure a group of operations

### Parameters
You can define route parameters using OpenAPI's parameters specification. Each parameter defined is used to validate
incoming data from the request. WeOS supports header, path, and query parameters (we don't support cookie parameters at
the time of writing). You can specify a parameter on a path or a specific operation within a path. Parameters that are
defined are accessible to middleware and controllers via the request context.

#### Route Level Parameters
```yaml
paths:
  /blogs:
    parameters:
        - in: query
          name: header
          schema:
            type: integer
        - in: query
          name: page
          schema:
            type: integer
```

#### Operation Level Parameters
```yaml
paths:
  /blogs:
    get:
        parameters:
            - in: query
              name: header
              schema:
                type: integer
            - in: query
              name: page
              schema:
                type: integer
```

### Middleware 
Middleware is reusable code that can be associated with a route. Middleware can add information to the request context
that can be used by the controller. To setup middleware use the `x-middleware` extension which allows for an array of 
middleware by name. The middlware needs to be regisered with the API before it can be used. WeOS provides standard middleware that you can use.

### Controllers
Controllers associated with the paths receive requests and execute commands or query data. To associate a Controller
with an endpoint, use the `x-controller` extension with the Controller name as a string.

#### Standard Controllers
To make it easier for you to get started, WeOS provides standard controllers for common data functionality. Controllers
are available for:

| Controller Name | Description                        | Conditions              | Required Parameters                            | Optional Parameters                                                                                                           |
|:----------------|:-----------------------------------|:------------------------|------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------|
| Create          | Create An Item                     | Schema associated with  | none                                           |                                                                                                                               |
| View            | Get the details of a specific item |                         | identifier (note this could be multiple parts) | use_entity_id (for getting item by entity id instead of by user defined id), sequence_no (get a specific version of the item) |
| List            | Get a collection of items          |                         | none                                           | page, limit, query, filters,                                                                                                  |
| CreateBatch     | Bulk create items                  |                         | none                                           |                                                                                                                               |
| Update          | Edit an item                       |                         | identifier (note this could be multiple parts) |                                                                                                                               |

Standard Controllers are automatically associated with an endpoint if a controller is not explicitly specified and the
path specification meets the conditions for one of the Standard Controllers. [Learn More About Controllers](/docs/concepts/controllers)

### Route Extensions
