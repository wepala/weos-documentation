---
title: Configure API info
weight: 10
description: Configure basic API info using the OpenAPI Info Object.
---

The [OpenAPI Info Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#infoObject) is used to
provide basic information about the API. This information can be displayed to users by using the health check standard controller.

```yaml
openapi: 3.0.3
info:
  title: Blog
  description: Blog example
  version: 1.0.0
```
