---
title: Projections
description: Map database content to data models.
---

In WeOS a projection is a data representation of an event stream. A developer can create a projection that takes events
as an input and stores it in a way that is easier for the application to use.

## Default Projection
A default projection is created using an Object Relationship Mapper (ORM) and the database configuration in `weos-config`.
The default projection supports a few databases including sqlite, Postgresql, Mysql and MariaDB. 
This projection will create a table for each schema in the specification and each entity will be a record in the
respective table.

## Custom Projections

There are a few reasons for creating custom projections including:

1. Using an unsupported database
2. Using events to generate data that should be re-created whenever the events are played (e.g. using events to generate QRcode images)

In order to create your own custom projection you must create a struct that implements `weos.Projection`

```go
//CustomProjection example of a custom projection
type CustomProjection struct {
}

func (c *CustomProjection) Migrate(ctx context.Context, builders map[string]dynamicstruct.Builder, deletedFields map[string][]string) error {
    return nil
}

func (c *CustomProjection) GetEventHandler() model.EventHandler {
    return func(ctx context.Context, event model.Event) error {
        switch event.Type {
            default:
                println("event handler hit")
        }
        return nil
    }
}

func (c *CustomProjection) GetContentEntity(ctx context.Context, entityFactory model.EntityFactory, weosID string) (*model.ContentEntity, error) {
//TODO implement me
panic("implement me")
}

func (c *CustomProjection) GetByKey(ctxt context.Context, entityFactory model.EntityFactory, identifiers map[string]interface{}) (map[string]interface{}, error) {
    entity := make(map[string]interface{})
    entity["title"] = "Foo"
    entity["description"] = "Bar"
    return entity, nil
}

func (c *CustomProjection) GetByEntityID(ctxt context.Context, entityFactory model.EntityFactory, id string) (map[string]interface{}, error) {
    entity := make(map[string]interface{})
    entity["title"] = "Foo"
    entity["description"] = "Bar"
    return entity, nil
}

func (c *CustomProjection) GetContentEntities(ctx context.Context, entityFactory model.EntityFactory, page int, limit int, query string, sortOptions map[string]string, filterOptions map[string]interface{}) ([]map[string]interface{}, int64, error) {
    entity := make(map[string]interface{})
    entity["title"] = "Foo"
    entity["description"] = "Bar"
    return []map[string]interface{}{entity}, 1, nil
}

func (c *CustomProjection) GetByProperties(ctxt context.Context, entityFactory model.EntityFactory, identifiers map[string]interface{}) ([]map[string]interface{}, error) {
    //TODO implement me
    panic("implement me")
}
```
!!! Tip
The event handler typically has a switch that handles different event types

To make the projection accessible via the open api spec it needs to be registered on the dependency container

```go
func main() {
    ctxt := context.Background()
    //instantiate weos with a reference to the OpenAPI specification
    api, err := controllers.New("api.yaml")
    if err != nil {
        log.Fatalln("error loading api config", err)
    }
    //register projection
    api.RegisterProjection("Default",CustomProjection)
    //initialize API so that standard middleware,controller,projections etc are registered
    err = api.Initialize(ctxt)
    if err != nil {
        log.Fatalln("error initializing api", err)
    }
    //start API 
    controllers.Serve("8681", api)
}

```

!!! Tip
Projection event handlers should be done in an idempotent way to accommodate replaying of events during a system recovery

### Default projection
The default projection can be replaced by registering the custom projection as the `Default` projection in the container
```go
//register middleware
api.RegisterProjection("Default",CustomProjection)
```

If you replace the default projection then you will also need to configure a default event store as well

### Adding Multiple Projections
Multiple projections can be associated with an operation by using the `x-projections` extension.
```yaml
/tasks:
  get:
    operationId: getTasks
    x-projections:
      - Custom
      - Default
```

The order of the projections matters, especially for endpoints that return data from the projection. In the case there are
multiple projections, WeOS will loop through each projection until it gets non nil result. If there really was no result
WeOS would have made a call to each configured projection.

[Complete Example](../examples/customizations/multiple_projections)