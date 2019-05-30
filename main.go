package main

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/savaki/swag"
	"github.com/savaki/swag/endpoint"
	"github.com/savaki/swag/swagger"
)

func handle(c *gin.Context) {
	io.WriteString(c.Writer, "Insert your code here")
}

type HTTPHeaderField struct {
	Type  string `json:type`
	Value string `json:value`
}

type HTTPResponse struct {
	Header	[] HTTPHeaderField `json:header,omitempty`
	Body string `json:body`
}

type Endpoint struct {
	Endpoint string `json:endpoint`
	Method string `json:method`
	Response HTTPResponse `json:response`
}

func main() {
	rest_post := endpoint.New("post", "/endpoints", "Add a new endpoint.",
		endpoint.Handler(handle),
		endpoint.Description("Add a new endpoint which you want to use. You can add response format"),
		endpoint.Body(Endpoint{}, "Endpoint object. You can set endpoint, method and response format.", true),
		endpoint.Response(http.StatusOK, "Success.", "Successfully added endpoint"),
	)
	rest_get := endpoint.New("get", "/endpoints", "Get registed endpoints.",
		endpoint.Handler(handle),
		endpoint.Response(http.StatusOK, []Endpoint{}, "successful operation"),
	)
	rest_delete := endpoint.New("delete", "/endpoints/{endpoint}", "delete the endpoint.",
		endpoint.Handler(handle),
		endpoint.Description("Delete the endpoint."),
		endpoint.Path("endpoint", "string", "endpoint.", true),
		endpoint.Response(http.StatusOK, "Success", "Successfully deleted endpoint"),
	)

	api := swag.New(
		swag.Endpoints(rest_post, rest_get, rest_delete),
		swag.Title("API server simulator"),
		swag.Version("0.0.1"),
		swag.Description("This is the simulator of API server. You can regist REST API endpoint with response."),
		swag.ContactEmail("kikikaikaikai@gmail.com"),
		swag.Tag("","", swag.TagURL("http://example.com/")),
	)

	router := gin.New()
	api.Walk(func(path string, endpoint *swagger.Endpoint) {
		h := endpoint.Handler.(func(c *gin.Context))
		path = swag.ColonPath(path)

		router.Handle(endpoint.Method, path, h)
	})

	enableCors := true
	router.GET("/swagger", gin.WrapH(api.Handler(enableCors)))

	http.ListenAndServe(":8080", router)
}
