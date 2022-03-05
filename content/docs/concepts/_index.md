---
title: Concepts
weight: 30
description: Understand core concepts.
---

When WeOS receives a request to an endpoint, it first runs any
pre-[middleware][middleware]. It then gets the route specified in your OpenAPI spec, followed by running any additional [middleware][middleware]. From here,
the associated [controller][controllers] takes over, which executes the
[command][commands] to handle the request type. A [projection][projections] is created
and WeOS responds by returning relevant data.

```mermaid
flowchart TD
request[/Request/] --> preMiddleware[Run pre-middleware] --> getRoute[Get route] -->
middleware[Run middleware] --> controller[Run controller] -->
command[Execute command] --> projection[Create projection] -->
respond([Return response])
```

[middleware]: /docs/concepts/middleware
[controllers]: /docs/concepts/controllers
[commands]: /docs/concepts/commands
[projections]: /docs/concepts/projections
