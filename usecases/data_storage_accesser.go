package usecases

import (
	"github.com/developer-kikikaikai/api_server_simulator/datatypes"
)

type DataStorageAccesser interface {
	GetAll() []datatypes.Endpoint
	Get(uri string, method string) datatypes.Endpoint
	Add(endpoint datatypes.Endpoint) error
	Update(endpoint datatypes.Endpoint) error
	Delete(endpoint datatypes.Endpoint) error
	IsVisuableAPIDefinition() bool
	GetAPIEndpoint() string
}
