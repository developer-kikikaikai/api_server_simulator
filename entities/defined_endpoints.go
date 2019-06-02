package entities

import (
	"errors"

	"github.com/developer-kikikaikai/api_server_simulator/datatypes"
	"github.com/developer-kikikaikai/api_server_simulator/repository"
)

func GetAll() []datatypes.Endpoint {
	return repository.GetDataStorageAccesser().GetAll()
}

func get(uri string, method string) (*datatypes.Endpoint, error) {
	endpoints := repository.GetDataStorageAccesser().GetAll()
	for _, endpoint_info := range endpoints {
		if (endpoint_info.Endpoint != uri) || (endpoint_info.Method != method) {
			continue
		}
		return &endpoint_info, nil
	}

	return nil, errors.New("No data related to uri and method")
}

func Add(endpoint *datatypes.Endpoint) error {
	return repository.GetDataStorageAccesser().Add(endpoint)
}

func Update(endpoint *datatypes.Endpoint) error {
	return repository.GetDataStorageAccesser().Update(endpoint)
}

func Delete(uri string, method string) error {
	return repository.GetDataStorageAccesser().Delete(uri, method)
}

func GetResponse(uri string, method string) (*datatypes.HTTPResponse, error) {
	endpoint_info, err := get(uri, method)
	if err != nil {
		return nil, err
	}

	return &endpoint_info.Response, nil
}
