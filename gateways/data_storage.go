package gateways

import (
	"errors"

	"github.com/developer-kikikaikai/api_server_simulator/datatypes"
)

type DataStorage struct {
}

func (this *DataStorage) GetAll() []datatypes.Endpoint {
	return nil
}
func (this *DataStorage) Get(uri string, method string) datatypes.Endpoint {
	return datatypes.Endpoint{}
}

func (this *DataStorage) Add(endpoint datatypes.Endpoint) error {
	return errors.New("Not implement yet")
}

func (this *DataStorage) Update(endpoint datatypes.Endpoint) error {
	return errors.New("Not implement yet")
}

func (this *DataStorage) Delete(endpoint datatypes.Endpoint) error {
	return errors.New("Not implement yet")
}

func (this *DataStorage) IsVisuableAPIDefinition() bool {
	return true
}

func (this *DataStorage) GetAPIEndpoint() string {
	return "/endpoints"
}
