package usecases

import (
	"github.com/developer-kikikaikai/api_server_simulator/datatypes"
)

type DefinedEndpointController interface {
	IsHandled(endpoint string, method string) bool
	GetResponse(endpoint string, method string) *datatypes.HTTPResponse
}
