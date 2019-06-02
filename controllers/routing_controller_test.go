package controllers

import (
	"testing"
)

func testGet(t *testing.T) {
	t.Fatal("failed test")
}

func testPost(t *testing.T) {
	t.Fatal("failed test")
}

func testUpdate(t *testing.T) {
	t.Fatal("failed test")
}

func testDelete(t *testing.T) {
	t.Fatal("failed test")
}

func testDefinedEndpoints(t *testing.T) {
	t.Fatal("failed test")
}

func startHTTPServer() {
}

func stopHTTPServer() {
}

func TestRouting(t *testing.T) {
	/*start routing by goroutin*/
	startHTTPServer()

	/*call routing*/
	testGet(t)
	testPost(t)
	testUpdate(t)
	testDelete(t)

	/*exit server*/
	stopHTTPServer()
}
