package ioriver

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testDomainName  = "ioriver-go-test"
	testServiceId   = "255e9621-15f6-49b3-af12-ea980206f5f6"
	testObjectId    = "56bef007-2d0b-4f10-974d-94ee714b44fd"
	testPathPattern = "/test_path/*"
	testUrl         = "https://www.example.com/api/test"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the API client being tested.
	client *IORiverClient

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// Cloudflare client configured to use test server
	client = NewClient("dummyapitoken")
	client.EndpointUrl = server.URL + "/"
}

func teardown() {
	server.Close()
}

func RunList[T interface{}](t *testing.T, lister func(client *IORiverClient) ([]T, error), path string,
	serverData string, expected []T) {

	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `%s`, serverData)
	}

	mux.HandleFunc(path, handler)

	actual, err := lister(client)
	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func RunServiceList[T interface{}](t *testing.T, lister func(client *IORiverClient, serviceId string) ([]T, error), path string, serviceId string,
	serverData string, expected []T) {

	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `%s`, serverData)
	}

	mux.HandleFunc(path, handler)

	actual, err := lister(client, serviceId)
	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func RunGet[T interface{}](t *testing.T, getter func(client *IORiverClient, id string) (*T, error), path string, id string,
	serverData string, expected *T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `%s`, serverData)
	}

	mux.HandleFunc(path, handler)

	actual, err := getter(client, id)
	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func RunServiceGet[T interface{}](t *testing.T, getter func(client *IORiverClient, serviceId string, id string) (*T, error),
	path string, serviceId string, id string, serverData string, expected *T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `%s`, serverData)
	}

	mux.HandleFunc(path, handler)

	actual, err := getter(client, serviceId, id)
	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func RunCreate[T interface{}](t *testing.T, creator func(client *IORiverClient, newObj T) (*T, error), path string, newObj T,
	serverData string, expected *T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `%s`, serverData)
	}

	mux.HandleFunc(path, handler)

	actual, err := creator(client, newObj)
	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}
