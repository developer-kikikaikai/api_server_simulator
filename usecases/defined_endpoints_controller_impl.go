package usecases

import (
	"github.com/developer-kikikaikai/api_server_simulator/datatypes"
	"github.com/developer-kikikaikai/api_server_simulator/entities"
)

type DefinedEndpointControllerImpl struct {
}

func (this *DefinedEndpointControllerImpl) GetDefinedResponse(endpoint string, method string) (*datatypes.HTTPResponse, error) {
	return entities.GetResponse(endpoint, method)
}
