package usecases

import (
	"../datatypes"
)

type DefinedEndpointController interface {
	IsHandled(endpoint string, method string) bool
	GetResponse(endpoint string, method string) *datatypes.HTTPResponse
}
