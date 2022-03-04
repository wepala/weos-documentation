---
title: WeOS configuration
description: >
    Configure WeOS with databases, event sources, and middleware.
---

WeOS is configured using the `x-weos-config` extension.

Configuration allows:
* connecting to databases
* connecting to event sources
* specifying active REST middleware

Here is an example:
```yaml
x-weos-config:
  event-source:
    - title: default
      driver: service
      endpoint: https://prod1.weos.sh/events/v1
    - title: event
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

### Middleware

Middleware is listed under the `middleware` key:
```yaml
x-weos-config:
  middleware:
    ...
```

Pre-built middleware is detailed [here][middleware].

[middleware]: /docs/concepts/middleware

### Databases

Database connections are established using the `database` key:
```yaml
x-weos-config:
  database:
    host: ${POSTGRES_HOST}
    database: ${POSTGRES_DB}
    username: ${POSTGRES_USER}
    password: ${POSTGRES_PASSWORD}
    port: ${POSTGRES_PORT}
```

### Event sources

Event sources are specified with the `event-source` key:
```yaml
x-weos-config:
  event-source:
    - title: default
      driver: service
      endpoint: https://prod1.weos.sh/events/v1
    - title: event
      driver: sqlite3
      database: test.db
```
