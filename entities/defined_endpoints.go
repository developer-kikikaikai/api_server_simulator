package entities

import (
	"errors"
	"fmt"

	"github.com/developer-kikikaikai/api_server_simulator/datatypes"
	"github.com/developer-kikikaikai/api_server_simulator/repository"
)

func GetAll() []datatypes.Endpoint {
	return repository.GetDataStorageAccesser().GetAll()
}

func get(uri string, method string) (*datatypes.Endpoint, error) {
	endpoints := repository.GetDataStorageAccesser().GetAll()
	fmt.Printf("check uri. Current data is %v\n", endpoints)
	fmt.Printf("uri %v, method %v\n", uri, method)
	for _, endpoint_info := range endpoints {
		fmt.Printf("check uri. Current data is %v\n", endpoint_info)
		if (endpoint_info.Endpoint != uri) || (endpoint_info.Method != method) {
			fmt.Printf("continue %v %v %v %v \n", endpoint_info.Endpoint, uri, endpoint_info.Method, method)
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

	fmt.Printf("GetResponse: %v\n", endpoint_info.Response)
	return &endpoint_info.Response, nil
}
