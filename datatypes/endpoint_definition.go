package datatypes

type HTTPHeaderField struct {
	Type  string `json:type`
	Value string `json:value`
}

type HTTPResponse struct {
	Header []HTTPHeaderField `json:header,omitempty`
	Status int               `json:status`
	Body   string            `json:body`
}

type Endpoint struct {
	Endpoint string       `json:endpoint`
	Method   string       `json:method`
	Response HTTPResponse `json:response`
}

type ExistingEndpoint struct {
	Endpoint string `json:endpoint`
	Method   string `json:method`
}
