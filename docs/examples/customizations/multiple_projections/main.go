package main

import (
	dynamicstruct "github.com/ompluscator/dynamic-struct"
	controllers "github.com/wepala/weos/controllers/rest"
	model "github.com/wepala/weos/model"
	"golang.org/x/net/context"
	"log"
)

//CustomProjection example of a custom projection
type CustomProjection struct {
}

func (c *CustomProjection) Migrate(ctx context.Context, builders map[string]dynamicstruct.Builder, deletedFields map[string][]string) error {
	return nil
}

func (c *CustomProjection) GetEventHandler() model.EventHandler {
	return func(ctx context.Context, event model.Event) error {
		println("event handler hit")
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

func main() {
	ctxt := context.Background()
	//instantiate weos with a reference to the OpenAPI specification
	api, err := controllers.New("api.yaml")
	if err != nil {
		log.Fatalln("error loading api config", err)
	}
	//register projection
	api.RegisterProjection("Custom", &CustomProjection{})
	//initialize API so that standard middleware,controller,projections etc are registered
	err = api.Initialize(ctxt)
	if err != nil {
		log.Fatalln("error initializing api", err)
	}
	//start API
	controllers.Serve("8681", api)
}
