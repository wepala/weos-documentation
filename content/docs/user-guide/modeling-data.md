---
title: Modeling data
weight: 20
description: Model data with OpenAPI schemas.
---

Data is modeled using [OpenAPI schemas][schemas] directly in your OpenAPI spec file.
For example, a blog API will have the concept of a "Blog" and a "Post":

[schemas]: https://swagger.io/docs/specification/data-models/

```yaml
 Blog:
      type: object
      properties:
        url:
          type: string
          format: uri
        title:
          type: string
        description:
          type: string
        status:
          type: string
          nullable: true
          enum:
            - null
            - unpublished
            - published
        image:
          type: string
          format: byte
        categories:
          type: array
          items:
            $ref: "#/components/schemas/Category"
        posts:
          type: array
          items:
            $ref: "#/components/schemas/Post"
        lastUpdated:
          type: string
          format: date-time
        created:
          type: string
          format: date-time
      required:
        - title
        - url
```

### Defining Properties
You can define properties using the [standard OpenAPI property syntax][data-types] and types. You can also use the "string" property
type along with the "format" attribute to specify additional types (e.g., date-time). See the property specification for
a complete list of property types and formats. 

| Data Type | Description                                       | OpenAPI Data Type | Format    | Default Value | 
|:----------|:--------------------------------------------------|:------------------|:----------|:--------------|
| Integer   |                                                   | integer           |           | nil           |
| Number    | Floating point Number                             | number            |           |               |
| Boolean   | true or false only. Truthy values are not allowed | boolean           |           |               |
| String    | string                                            | string            |           |               |
| Date Time |                                                   | string            | date-time |               |
| Array     |                                                   | array             |           |               |
| Object    |                                                   |                   |           |               |

[data-types]: https://swagger.io/docs/specification/data-models/data-types/

### Setting Identifiers
You can use one (or more) of the properties you defined as an identifier for the schema by using the `x-identifier`
attribute. The `x-identifier` attribute is a list of properties that you want to use to identify an instance of the schema
uniquely. 

```yaml
Blog:
      type: object
      properties:
        id:
          type: string
        title:
          type: string
        description:
          type: string
      x-identifier:
        - id
        - title
```
Each schema must have an identifier, so if one is not explicitly defined, WeOS will automatically
add a property "id" to the schema.

### Validation
To specify basic business rules, you can use the standard OpenAPI "required" attribute on a Content Type to indicate
which properties are required. You can also use the "pattern" attribute on a specific property to specify a RegEx to use
for validation.
