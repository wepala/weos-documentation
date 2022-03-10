---
title: WEOS CLI
weight: 10
description: Running the WeOS CLI
---
Using the CLI is by far the easiest starting point. [Download][download] and run the CLI by pointing it to your OpenAPI spec.
By default, the API will run on port `8681` (you can configure this using the `--port` switch), and it will try to use a
specification file named `api.yaml` (you can specify this using the `--spec` switch).

The WeOS CLI is essentially a server that uses the OpenAPI specification for configuration. We chose to build the
server with Go because we wanted to make the server extensible, easy to deploy and maintain with no serverside runtime required.

[download]: https://github.com/wepala/weos/releases