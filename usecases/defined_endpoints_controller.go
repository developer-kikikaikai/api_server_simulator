package usecases

import (
	"github.com/developer-kikikaikai/api_server_simulator/datatypes"
)

type DefinedEndpointController interface {
	GetDefinedResponse(endpoint string, method string) (*datatypes.HTTPResponse, error)
}
