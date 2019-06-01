package controllers

import (
	"fmt"
	"log"
	"net/http"

	"../datatypes"
	"../factory"
	"github.com/gin-gonic/gin"
	"github.com/savaki/swag"
	"github.com/savaki/swag/endpoint"
	"github.com/savaki/swag/swagger"
)

func createSwagAPI() *swagger.API {
	uri := factory.GetDataStorageAccesser().GetAPIEndpoint()
	fmt.Printf("%+v\n", uri)
	/*add endpoint /endpints to regist some request */
	rest_get := endpoint.New("get", uri, "Get registed endpoints.",
		endpoint.Handler(GETEndpointsHandler),
		endpoint.Response(http.StatusOK, []datatypes.Endpoint{}, "successful operation"),
	)
	rest_post := endpoint.New("post", uri, "Add a new endpoint.",
		endpoint.Handler(POSTEndpointsHandler),
		endpoint.Description("Add a new endpoint which you want to use. You can add response format"),
		endpoint.Body(datatypes.Endpoint{}, "Endpoint object. You can set endpoint, method and response format.", true),
		endpoint.Response(http.StatusOK, "Success.", "Successfully added endpoint"),
	)
	rest_put := endpoint.New("put", uri, "Update the endpoint.",
		endpoint.Handler(PUTEndpointsHandler),
		endpoint.Description("Update the endpoint which you want to use."),
		endpoint.Path("endpoint", "string", "endpoint.", true),
		endpoint.Body(datatypes.ExistingEndpoint{}, "Endpoint object. You can set endpoint, method and response format.", true),
		endpoint.Response(http.StatusOK, "Success.", "Successfully added endpoint"),
	)
	rest_delete := endpoint.New("delete", uri, "delete the endpoint.",
		endpoint.Handler(DELETEEndpointsHandler),
		endpoint.Description("Delete the endpoint."),
		endpoint.Body(datatypes.ExistingEndpoint{}, "Endpoint object. You can set endpoint, method and response format.", true),
		endpoint.Response(http.StatusOK, "Success", "Successfully deleted endpoint"),
	)

	return swag.New(
		swag.Endpoints(rest_post, rest_put, rest_get, rest_delete),
		swag.Title("API server simulator"),
		swag.Version("0.0.1"),
		swag.Description("This is the simulator of API server. You can regist REST API endpoint with response."),
		swag.ContactEmail("kikikaikaikai@gmail.com"),
		swag.Tag("", "", swag.TagURL("http://example.com/")),
	)
}

func HandleDefinedEndpoints() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("before logic")
		c.Next()
	}
}

func CreateRouter() *gin.Engine {
	//create swagger api instance
	api := createSwagAPI()

	router := gin.New()
	api.Walk(func(path string, endpoint *swagger.Endpoint) {
		h := endpoint.Handler.(func(c *gin.Context))
		path = swag.ColonPath(path)
		router.Handle(endpoint.Method, path, h)
	})

	//set middleware handler to catch any event
	router.Use(HandleDefinedEndpoints())

	//enable swagger endpoint if enable option
	if factory.GetDataStorageAccesser().IsVisuableAPIDefinition() {
		router.GET("/swagger", gin.WrapH(api.Handler(true)))
	}
	return router
}
