package custom_middleware

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	controllers "github.com/wepala/weos/controllers/rest"
	model "github.com/wepala/weos/model"
	projections "github.com/wepala/weos/projections"
	"golang.org/x/net/context"
	"log"
)

func CustomMiddleware(api *controllers.RESTAPI, projection projections.Projection, commandDispatcher model.CommandDispatcher, eventSource model.EventRepository, entityFactory model.EntityFactory, path *openapi3.PathItem, operation *openapi3.Operation) echo.MiddlewareFunc {
	//TODO code you want to execute once can go here e.g. parsing information from OpenAPI spec
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctxt echo.Context) error {
			//TODO code that you want to execute per request goes here.
			//variables define in the OpenAPI spec will be available in the request context ctxt.Request().Context()
			return nil
		}
	}
}

func main() {
	ctxt := context.Background()
	//instantiate weos with a reference to the OpenAPI specification
	api, err := controllers.New("api.yaml")
	if err != nil {
		log.Fatalln("error loading api config", err)
	}
	//register middleware
	api.RegisterMiddleware("Custom", CustomMiddleware)
	//initialize API so that standard middleware,controller,projections etc are registered
	err = api.Initialize(ctxt)
	if err != nil {
		log.Fatalln("error initializing api", err)
	}
	//start API
	controllers.Serve("8681", api)
}
