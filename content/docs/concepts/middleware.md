---
title: Middleware
---

Middleware is reusable code that can be associated with a route.

Built-in middleware includes:
- **RequestID**: generate request IDs
- **Recover**: recover from panics and handle appropriately
- **Logger**: log information about HTTP requests
- **ZapLogger**: log using structured, leveled logging
- **CSVUpload**: parse CSV files

Middleware is configured in your OpenAPI spec with the `x-weos-config`
extension:

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
