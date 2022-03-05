---
title: Controllers
description: Control handling of endpoint requests.
---

[projections]: /docs/concepts/projections

Controllers map requests to models as well as retrieve information through projections. WeOS provides standard controllers for everyday actions, e.g., Create, Read, Update, Delete (CRUD).

WeOS automatically binds controllers to endpoints that don't have one already. The controller is automatically attached based on the HTTP method, request body, and response info.

## Create
The Create controller will create an item of the content type associated with the controller's endpoint.

Example OpenAPI usage:
```yaml
paths:
  /posts/{postId}:
     put:
      parameters:
        - in: path
          name: postId
          schema:
            type: string
          required: true
      summary: Update post
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/Post"
```

## Read
There are two controllers for reading data from WeOS.
- **Get**: Get a single item
- **List**: Get a collection of items

### Get
The Get controller retrieves an item using the identifier of the associated content
type. The developer defines an endpoint with the identifier declared as a parameter to
use the Get controller. The Get controller retrieves the identifier from the request
context and uses the [projection][projections] to get the information from the database.

Example OpenAPI usage:
```yaml
paths:
  /posts/{postId}:
    get:
      parameters:
        - in: path
          name: postId
          schema:
            type: string
          required: true
      summary: Get blog post by id
      responses:
        200:
          description: Get blog post information
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Post"
```

### List
The List controller returns a paginated collection of items based on filters and sorts. The controller uses the filters
declared in the context to retrieve information using a [projection][projections].

Example OpenAPI usage:
```yaml
paths:
  /posts/:
    get:
      operationId: Get Posts
      summary: Get a blog's list of posts
      parameters:
        - in: query
          name: q
          schema:
            type: string
          required: false
          description: query string
      responses:
        200:
          description: List of blog posts
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
                      $ref: "#/components/schemas/Post"
```

## Update
The Update controller will update an item using the identifier of the associated
content type.

Example OpenAPI usage:
```yaml
paths:
  /posts/{postId}:
    put:
      parameters:
        - in: path
          name: postId
          schema:
            type: string
          required: true
      summary: Update post
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/Post"
      responses:
        200:
          description: Get blog post information
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Post"
```
## Delete
The Delete controller will delete an item using the identifier of the associated
content type.

Example OpenAPI usage:
```yaml
paths:
  /posts/{postId}:
    delete:
      parameters:
        - in: path
          name: postId
          schema:
            type: string
          required: true
      summary: Delete post
      responses:
        200:
          description: Delete post
```
