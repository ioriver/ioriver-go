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

func TestCreateServiceAttachedFile(t *testing.T) {
	setup()
	defer teardown()

	serviceID := "255e9621-15f6-49b3-af12-ea980206f5f6"
	mux.HandleFunc("/v1/services/"+serviceID+"/attached_files/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "token dummyapitoken", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("content-type"))

		var payload map[string]string
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, map[string]string{
			"service":   serviceID,
			"file_type": "HTML",
			"contents":  "<html>blocked</html>",
		}, payload)

		w.Header().Set("content-type", "application/json")
		_, _ = w.Write([]byte(`{"id":"3260cb66-ebbc-4b62-b647-3f3fbe0100da","service":"` + serviceID + `","file_type":"HTML"}`))
	})

	file, err := client.CreateServiceAttachedFile(serviceID, "HTML", "<html>blocked</html>")
	if assert.NoError(t, err) {
		assert.Equal(t, "3260cb66-ebbc-4b62-b647-3f3fbe0100da", file.ID)
		assert.Equal(t, serviceID, file.Service)
		assert.Equal(t, "HTML", file.FileType)
	}
}
