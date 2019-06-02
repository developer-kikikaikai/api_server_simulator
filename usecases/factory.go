package usecases

//define all interface
var definedEndpointController = DefinedEndpointControllerImpl{}
var apiController = APIControllerImpl{}

func GetDefinedEndpointController() DefinedEndpointController {
	return &definedEndpointController
}

func GetAPIController() APIController {
	return &apiController
}
