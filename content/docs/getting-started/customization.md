---
title: Customization
weight: 20
description: Build a custom WeOS binary.
---

If you know a bit about Go, you can use WeOS as a library for creating REST API using OpenAPI (as opposed to using the weos-cli). In WeOS middleware,
controllers and initializers are simply functions you can create and store in the dependency container. Dependencies are
registered can be referenced in the OpenAPI specification. Projections, event stores, and command dispatchers can all be
customized.

## Starting A Customized API

To get started with building your own custom Go application using WeOS

1. Install `go get github.com/wepala/weos`
2. Setup api
```go
package main

import (
	weos "github.com/wepala/weos/controllers/rest"
	"golang.org/x/net/context"
	"log"
)

func main() {
	ctxt := context.Background()
	//instantiate weos with a reference to the OpenAPI specification
	api, err := weos.New("api.yaml")
	if err != nil {
		log.Fatalln("error loading api config", err)
	}
   //initialize API so that standard middleware,controller,projections etc are registered
	err = api.Initialize(ctxt)
	if err != nil {
		log.Fatalln("error initializing api", err)
	}
	//start API 
	weos.Serve("8681", api)
}
```
3. Compile the app using `go build`
4. Run api by calling the app that was compiled in the previous step


## Customizations
1. [Create Custom Middleware](/docs/user-guide/middleware#adding-custom-middleware)
