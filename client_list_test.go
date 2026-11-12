package ioriver

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListFunctionCompatibility(t *testing.T) {
	setup()
	defer teardown()

	var list func(*IORiverClient, string) ([]Service, error) = List[Service]
	mux.HandleFunc("/v1/services/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Empty(t, r.URL.RawQuery)
		_, _ = w.Write([]byte(`[]`))
	})

	services, err := list(client, servicesBasePath)
	if !assert.NoError(t, err) {
		return
	}
	assert.Empty(t, services)
}

func TestListWithQuery(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/services/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testTagId, r.URL.Query().Get("tag_ids"))
		assert.Equal(t, "prod", r.URL.Query().Get("name"))
		_, _ = w.Write([]byte(`[]`))
	})

	services, err := ListWithQuery[Service](client, servicesBasePath, "tag_ids="+testTagId, "name=prod")
	if !assert.NoError(t, err) {
		return
	}
	assert.Empty(t, services)
}
