---
title: OpenAPI usage
weight: 10
---

WeOS takes a design-first approach using [OpenAPI specifications](https://www.openapis.org/). WeOS uses the OpenAPI specification to set up routes
and automatically associate controllers. WeOS attempts to allow developers to get a lot done with "vanilla" OpenAPI,
although you can use OpenAPI extensions to provide additional customizations and configurations.

## Configuring basic api info 
The [OpenAPI Info Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#infoObject) is used to
provide basic information about the api. This information can be displayed to users by using the health check standard controller
```yaml
openapi: 3.0.3
info:
  title: Blog
  description: Blog example
  version: 1.0.0
```

TODO: Working with models and relationships
TODO: Setting up a custom API
TODO: Setting up a custom projection
TODO: Using your own controller
TODO: Creating a custom command