package entities

import (
	"errors"

	"github.com/developer-kikikaikai/api_server_simulator/datatypes"
	"github.com/developer-kikikaikai/api_server_simulator/factory"
)

type DefinedEndpoints struct {
}

var definedEndpoints = DefinedEndpoints{}

func (this *DefinedEndpoints) GetAll() []datatypes.Endpoint {
	return factory.GetDataStorageAccesser().GetAll()
}

func (this *DefinedEndpoints) Get(uri string, method string) (datatypes.Endpoint, error) {
	if factory.GetDefinedEndpointController().IsHandled(uri, method) {
		return factory.GetDataStorageAccesser().Get(uri, method), nil
	} else {
		return _, errors.New("No data related to uri and method")
	}
}

func (this *DefinedEndpoints) Add(endpoint Endpoint) error {
	return factory.GetDataStorageAccesser().Add(endpoint)
}

func (this *DefinedEndpoints) Update(endpoint Endpoint) error {
	return factory.GetDataStorageAccesser().Update(endpoint)
}

func (this *DefinedEndpoints) Delete(endpoint Endpoint) error {
	return factory.GetDataStorageAccesser().Delete(endpoint)
}

func (this *DefinedEndpoints) GetResponse(uri string, method string) (*datatypes.HTTPResponse, error) {
	api := factory.GetDefinedEndpointController()
	if api.IsHandled(uri, method) {
		return api.GetResponse(uri, method)
	} else {
		return errors.New("No data related to uri and method")
	}
}
