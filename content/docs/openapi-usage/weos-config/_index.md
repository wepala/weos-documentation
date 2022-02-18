---
title: WeOS config
---

WeOS is configured using the `x-weos-config` [OpenAPI extension][extensions].

Configuration allows:
* connecting to databases
* connecting to event sources
* specifying active REST middleware

```yaml
x-weos-config:
  event-source:
    - title: default
      driver: service
      endpoint: https://prod1.weos.sh/events/v1
    - title: event
      driver: sqlite3
      database: test.db
  databases:
    - title: default
      driver: sqlite3
      database: test.db
  database:
    driver: sqlite3
    database: test.db
  rest:
    middleware:
      - RequestID
      - Recover
      - ZapLogger
```

TODO: Document available keys and values

[extensions]: https://swagger.io/docs/specification/openapi-extensions/