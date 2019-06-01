package usecases

import (
	"github.com/developer-kikikaikai/api_server_simulator/datatypes"
)

type DefinedEndpointControllerImpl struct {
}

func (this *DefinedEndpointControllerImpl) IsHandled(endpoint string, method string) bool {
	return false
}

func (this *DefinedEndpointControllerImpl) GetResponse(endpoint string, method string) *datatypes.HTTPResponse {
	return nil
}
