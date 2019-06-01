package entities

import (
	"errors"

	"../datatypes"
)

type DefinedEndpoints struct {
}

var definedEndpoints = DefinedEndpoints{}

func (this *DefinedEndpoints) GetAll() []datatypes.Endpoint {
	return nil
}

func (this *DefinedEndpoints) Get(uri string, method string) (datatypes.Endpoint, error) {
	return nil, errors.New("Not implement yet")
}
func (this *DefinedEndpoints) Add(endpoint Endpoint) error {
	return errors.New("Not implement yet")
}
func (this *DefinedEndpoints) Update(endpoint Endpoint) error {
	return errors.New("Not implement yet")
}
func (this *DefinedEndpoints) Delete(endpoint Endpoint) error {
	return errors.New("Not implement yet")
}
func (this *DefinedEndpoints) GetResponse(uri string, method string) (datatypes.HTTPResponse, error) {
	return nil, errors.New("Not implement yet")
}
