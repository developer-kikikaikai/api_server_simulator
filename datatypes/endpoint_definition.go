package datatypes

/*data format for Configuration*/
type HTTPSConfiguration struct {
	Port            int    `json:port`
	CertificatePath string `json:certificate`
	PrivatekeyPath  string `json:privatekeyPath`
}

/*data format for API request*/
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

const PostSuccess = "Success to add Endpoint"
const PutSuccess = "Success to update Endpoint"
const DeleteSuccess = "Success to delete Endpoint"
const ParseFailed = "Failed to parse request body"
const FailedToAdd = "Failed to add request Endpoint"
const FailedToUpdate = "Failed to update request Endpoint"
const FailedToDelete = "Failed to delete request Endpoint"

func IsEqualEndpointKey(src *Endpoint, dist *Endpoint) bool {
	return (src.Endpoint == dist.Endpoint) && (src.Method == dist.Method)
}

func isInHeaders(src *HTTPHeaderField, dist []HTTPHeaderField) bool {
	for _, dist_header := range dist {
		if src.Type == dist_header.Type && src.Value == dist_header.Value {
			return true
		}
	}

	return false
}

func IsEqualEndpoint(src *Endpoint, dist *Endpoint) bool {
	//check 1st rank
	if !IsEqualEndpointKey(src, dist) {
		return false
	}

	//check response data
	if src.Response.Status != dist.Response.Status || src.Response.Body != dist.Response.Body {
		return false
	}

	//check response header
	if len(src.Response.Header) != len(dist.Response.Header) {
		return false
	}

	for _, src_header := range src.Response.Header {
		if !isInHeaders(&src_header, dist.Response.Header) {
			return false
		}
	}
	return true
}
