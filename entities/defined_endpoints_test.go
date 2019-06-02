package entities

import (
	"net/http"
	"testing"

	"github.com/developer-kikikaikai/api_server_simulator/datatypes"
)

/*same as data*/
func testClearingData(t *testing.T) {
	endpoints := GetAll()
	for _, endpoint := range endpoints {
		Delete(endpoint.Endpoint, endpoint.Method)
	}

	endpoints = GetAll()
	if len(endpoints) != 0 {
		//Failed, but not stop running
		t.Error("Failed to Delete all, value is NULL")
	}
}

func TestGetAllAndAdd(t *testing.T) {
	testClearingData(t)

	/*No HTTP Header*/
	new_endpoint := datatypes.Endpoint{}
	new_endpoint.Endpoint = "test"
	new_endpoint.Method = "POST"
	new_endpoint.Response.Status = http.StatusBadRequest
	new_endpoint.Response.Body = "{\"result\":\"value\"}"
	if err := Add(&new_endpoint); err != nil {
		t.Logf("result: %v\n", err)
		t.Fatal("Failed to add data")
	}

	endpoints := GetAll()
	if len(endpoints) != 1 {
		//Stop testing
		t.Fatal("Failed to get new adding endpoint")
	}

	if !datatypes.IsEqualEndpoint(&endpoints[0], &new_endpoint) {
		t.Logf("result: %v and %v\n", endpoints[0], new_endpoint)
		t.Fatal("Failed to add data, get different endpoint")
	}

	/*Add HTTP Header*/
	new_endpoint2 := datatypes.Endpoint{}
	new_endpoint2.Endpoint = "test2"
	new_endpoint2.Method = "GET"
	new_endpoint2.Response.Status = http.StatusOK
	new_endpoint2.Response.Body = "[1]"
	header_field := datatypes.HTTPHeaderField{"Cache-Control", "no-cache "}
	new_endpoint2.Response.Header = append(new_endpoint2.Response.Header, header_field)
	header_field = datatypes.HTTPHeaderField{"Connection", "Keep-Alive"}
	new_endpoint2.Response.Header = append(new_endpoint2.Response.Header, header_field)

	if err := Add(&new_endpoint2); err != nil {
		t.Logf("result: %v\n", err)
		t.Fatal("Failed to add data with HTTP Header Field")
	}

	endpoints = GetAll()
	if endpoints == nil || len(endpoints) != 2 {
		//Stop testing
		t.Fatal("Failed to get re-adding endpoint")
	}

	if !datatypes.IsEqualEndpoint(&endpoints[0], &new_endpoint) && !datatypes.IsEqualEndpoint(&endpoints[1], &new_endpoint) {
		t.Logf("result: %v and %v\n", endpoints[0], endpoints[0])
		t.Fatal("Failed to add data, get different endpoint")
	}

	if !datatypes.IsEqualEndpoint(&endpoints[0], &new_endpoint2) && !datatypes.IsEqualEndpoint(&endpoints[1], &new_endpoint2) {
		t.Logf("result: %v and %v\n", endpoints[0], endpoints[0])
		t.Fatal("Failed to add data, get different endpoint")
	}

	/*Add same data*/
	if err := Add(&new_endpoint2); err == nil {
		t.Fatal("Failed to check same data.")
	}

	testClearingData(t)
}

func TestUpdate(t *testing.T) {
	testClearingData(t)

	/*No HTTP Header*/
	new_endpoint := datatypes.Endpoint{}
	new_endpoint.Endpoint = "test"
	new_endpoint.Method = "POST"
	new_endpoint.Response.Status = http.StatusBadRequest
	new_endpoint.Response.Body = "{\"result\":\"value\"}"
	if err := Add(&new_endpoint); err != nil {
		t.Logf("result: %v\n", err)
		t.Fatal("Failed to add data")
	}

	//Update
	header_field := datatypes.HTTPHeaderField{"Cache-Control", "no-cache "}
	new_endpoint.Response.Header = append(new_endpoint.Response.Header, header_field)
	header_field = datatypes.HTTPHeaderField{"Connection", "Keep-Alive"}
	new_endpoint.Response.Header = append(new_endpoint.Response.Header, header_field)
	if err := Update(&new_endpoint); err != nil {
		t.Logf("result: %v\n", err)
		t.Fatal("Failed to update data")
	}

	endpoints := GetAll()
	if len(endpoints) != 1 {
		t.Logf("endpoints: %v\n", endpoints)
		t.Fatal("Data is not updated, data is added or removed.")
	}

	if !datatypes.IsEqualEndpoint(&endpoints[0], &new_endpoint) {
		t.Logf("endpoints: %v\n", endpoints)
		t.Fatal("Data is not updated, some part is not overrided.")
	}

	//Change endpoint
	new_endpoint.Endpoint = "test1"
	if err := Update(&new_endpoint); err == nil {
		t.Fatal("Failed to check different endpoint updating\n")
	}

	//Change key
	new_endpoint.Endpoint = "test"
	new_endpoint.Method = "GET"
	if err := Update(&new_endpoint); err == nil {
		t.Fatal("Failed to check different method updating\n")
	}

	//check current data
	endpoints2 := GetAll()
	if len(endpoints2) != 1 {
		t.Logf("endpoints: %v\n", endpoints2)
		t.Fatal("Data is inserted or removed when reject different key updating.")
	}

	if !datatypes.IsEqualEndpoint(&endpoints[0], &endpoints2[0]) {
		t.Logf("endpoints: %v\n", endpoints)
		t.Logf("endpoints2: %v\n", endpoints2)
		t.Fatal("Data is updated when reject different key updating.")
	}

	testClearingData(t)
}

func TestDelete(t *testing.T) {
	testClearingData(t)

	/*No HTTP Header*/
	new_endpoint := datatypes.Endpoint{}
	new_endpoint.Endpoint = "test"
	new_endpoint.Method = "POST"
	new_endpoint.Response.Status = http.StatusBadRequest
	new_endpoint.Response.Body = "{\"result\":\"value\"}"
	if err := Add(&new_endpoint); err != nil {
		t.Logf("result: %v\n", err)
		t.Fatal("Failed to add data")
	}

	//delete no key
	if err := Delete(new_endpoint.Endpoint+"ss", new_endpoint.Method); err == nil {
		t.Fatal("Failed to check invalid Endpoint")
	}
	if err := Delete(new_endpoint.Endpoint, new_endpoint.Method+"ss"); err == nil {
		t.Fatal("Failed to check invalid Method")
	}

	//check current data
	endpoints := GetAll()
	if len(endpoints) != 1 {
		t.Logf("endpoints: %v\n", endpoints)
		t.Fatal("Data is removed when reject different key deleting.")
	}

	if !datatypes.IsEqualEndpoint(&new_endpoint, &endpoints[0]) {
		t.Logf("endpoints: %v\n", new_endpoint)
		t.Logf("endpoints2: %v\n", endpoints[0])
		t.Fatal("Data is updated when reject different key deleting.")
	}

	if err := Delete(new_endpoint.Endpoint, new_endpoint.Method); err != nil {
		t.Logf("result: %v\n", err)
		t.Fatal("Failed to delete")
	}

	endpoints = GetAll()
	if len(endpoints) != 0 {
		t.Fatal("Failed to delete")
	}
}

func TestGetResonse(t *testing.T) {
	testClearingData(t)
	/*Add HTTP Header*/
	new_endpoint2 := datatypes.Endpoint{}
	new_endpoint2.Endpoint = "test2"
	new_endpoint2.Method = "GET"
	new_endpoint2.Response.Status = http.StatusOK
	new_endpoint2.Response.Body = "[1]"
	header_field := datatypes.HTTPHeaderField{"Cache-Control", "no-cache "}
	new_endpoint2.Response.Header = append(new_endpoint2.Response.Header, header_field)
	header_field = datatypes.HTTPHeaderField{"Connection", "Keep-Alive"}
	new_endpoint2.Response.Header = append(new_endpoint2.Response.Header, header_field)

	if err := Add(&new_endpoint2); err != nil {
		t.Logf("result: %v\n", err)
		t.Fatal("Failed to add data with HTTP Header Field")
	}

	response, err := GetResponse(new_endpoint2.Endpoint, new_endpoint2.Method)
	if err != nil {
		t.Logf("result: %v\n", err)
		t.Fatal("Failed to get result")
	}

	new_endpoint := datatypes.Endpoint{new_endpoint2.Endpoint, new_endpoint2.Method, *response}
	if !datatypes.IsEqualEndpoint(&new_endpoint2, &new_endpoint) {
		t.Logf("endpoints: %v\n", new_endpoint)
		t.Logf("endpoints2: %v\n", new_endpoint2)
		t.Fatal("Failed to get correct response data")
	}

	/*check error validation*/
	_, err = GetResponse(new_endpoint2.Endpoint+"data", new_endpoint2.Method)
	if err == nil {
		t.Logf("result: %v\n", err)
		t.Fatal("Failed to check endpoint filter")
	}

	_, err = GetResponse(new_endpoint2.Endpoint, new_endpoint2.Method+"data")
	if err == nil {
		t.Logf("result: %v\n", err)
		t.Fatal("Failed to check method filter")
	}

	testClearingData(t)
}
