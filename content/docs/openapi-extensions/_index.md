---
title: OpenAPI extensions
weight: 40
description: Learn about WeOS-specific OpenAPI extensions.
---

WeOS provides [OpenAPI extensions][extensions] to describe its extra functionality.

These include:
- [`x-weos-config`][x-weos-config]: Configure WeOS
- `x-context-name`: Alias parameter name to a different name in the context
- `x-middleware`: Add middleware
- `x-controller`: Set controller
- `x-remove`: Mark a field for removal
- `x-copy`: Copy a field's value into another field
- `x-identifier`: Set the identifier
- `x-alias`: Alias parameter name to a different name in the controller
- `x-schema`: Specify the content type instead of the request body
- `x-projection`: Set a custom projection
- `x-command-dispatcher`: Set a custom command dispatcher
- `x-event-source`: Set a custom event source

[extensions]: https://swagger.io/docs/specification/openapi-extensions/
[x-weos-config]: /docs/openapi-extensions/x-weos-config
