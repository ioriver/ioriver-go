package ioriver

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListServicesByTagId(t *testing.T) {
	for _, tc := range []struct {
		name string
		id   string
		body string
		want []Service
	}{
		{"matching services", testTagId, `[{"id":"` + testServiceId + `","name":"service"}]`, []Service{{Id: testServiceId, Name: "service"}}},
		{"no matches", testTagId, `[]`, []Service{}},
		{"query escaping", "id&name=other", `[]`, []Service{}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			setup()
			defer teardown()

			mux.HandleFunc("/v1/services/", func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, tc.id, r.URL.Query().Get("tags"))
				assert.Len(t, r.URL.Query(), 1)
				_, _ = w.Write([]byte(tc.body))
			})

			services, err := client.ListServicesByTagId(tc.id)
			if !assert.NoError(t, err) {
				return
			}
			assert.Equal(t, tc.want, services)
		})
	}
}
func TestAddServiceTags(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/services/"+testServiceId+"/add_tags/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		var payload map[string][]map[string]string
		if !assert.NoError(t, json.NewDecoder(r.Body).Decode(&payload)) {
			return
		}
		assert.Equal(t, map[string][]map[string]string{
			"tags": {{"name": "prod"}, {"name": "edge"}},
		}, payload)
		_, _ = w.Write([]byte(`[{"tag_name":"prod","created":false},{"tag_name":"edge","created":true}]`))
	})

	results, err := client.AddServiceTags(testServiceId, []string{"prod", "edge"})
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, []ServiceTagResult{{TagName: "prod"}, {TagName: "edge", Created: true}}, results)
}

func TestRemoveServiceTag(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/services/"+testServiceId+"/remove_tag/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		var payload map[string]string
		if !assert.NoError(t, json.NewDecoder(r.Body).Decode(&payload)) {
			return
		}
		assert.Equal(t, map[string]string{"tag_name": "prod"}, payload)
		_, _ = w.Write([]byte(`{"tag_name":"prod","removed":true}`))
	})

	result, err := client.RemoveServiceTag(testServiceId, "prod")
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, &ServiceTagRemoval{TagName: "prod", Removed: true}, result)
}

func TestServiceTagAPIErrors(t *testing.T) {
	for _, tc := range []struct {
		name string
		call func(*IORiverClient) error
	}{
		{"filter", func(c *IORiverClient) error { _, err := c.ListServicesByTagId(testTagId); return err }},
		{"add", func(c *IORiverClient) error { _, err := c.AddServiceTags(testServiceId, []string{"prod"}); return err }},
		{"remove", func(c *IORiverClient) error { _, err := c.RemoveServiceTag(testServiceId, "prod"); return err }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			setup()
			defer teardown()

			mux.HandleFunc("/v1/services/", func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "service not found", http.StatusNotFound)
			})
			err := tc.call(client)
			if !assert.Error(t, err) {
				return
			}
			apiErr, ok := err.(*APIError)
			if !assert.True(t, ok, "expected APIError, got %T", err) {
				return
			}
			assert.Equal(t, http.StatusNotFound, apiErr.StatusCode)
		})
	}
}
