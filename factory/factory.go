package factory

import (
	"github.com/developer-kikikaikai/api_server_simulator/gateways"
	"github.com/developer-kikikaikai/api_server_simulator/usecases"
)

//define all interface
var definedEndpointController = usecases.DefinedEndpointControllerImpl{}
var apiController = usecases.APIControllerImpl{}
var dataStorage = gateways.DataStorage{}

func GetDefinedEndpointController() usecases.DefinedEndpointController {
	return &definedEndpointController
}

func GetAPIController() usecases.APIController {
	return &apiController
}

func GetDataStorageAccesser() usecases.DataStorageAccesser {
	return &dataStorage
}
