package ioriver

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testTagId = "826c5d31-03e1-4ad0-9e7a-19d6b12a7301"

const serverTagData = `{"id":"` + testTagId + `","name":"prod","account":"account-id","color":"#FF0000","services_count":2}`

var expectedTag = Tag{
	Id:            testTagId,
	Name:          "prod",
	Account:       "account-id",
	Color:         "#FF0000",
	ServicesCount: 2,
}

func TestGetTag(t *testing.T) {
	RunGet(t, (*IORiverClient).GetTag, "/tags/"+testTagId+"/", testTagId, serverTagData, &expectedTag)
}

func TestListTags(t *testing.T) {
	RunList(t, (*IORiverClient).ListTags, "/tags/", "["+serverTagData+"]", []Tag{expectedTag})
}

func TestWriteTag(t *testing.T) {
	for _, tc := range []struct {
		name   string
		method string
		path   string
		tag    Tag
		write  func(*IORiverClient, Tag) (*Tag, error)
	}{
		{"create", http.MethodPost, "/v1/tags/", Tag{Name: "prod", Color: "#FF0000"}, (*IORiverClient).CreateTag},
		{"update", http.MethodPut, "/v1/tags/" + testTagId + "/", expectedTag, (*IORiverClient).UpdateTag},
		{"default color", http.MethodPost, "/v1/tags/", Tag{Name: "prod"}, (*IORiverClient).CreateTag},
	} {
		t.Run(tc.name, func(t *testing.T) {
			setup()
			defer teardown()

			mux.HandleFunc(tc.path, func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, tc.method, r.Method)
				assert.Equal(t, "token dummyapitoken", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
				var payload map[string]interface{}
				if !assert.NoError(t, json.NewDecoder(r.Body).Decode(&payload)) {
					return
				}
				assert.Equal(t, "prod", payload["name"])
				if tc.tag.Color == "" {
					assert.NotContains(t, payload, "color")
				} else {
					assert.Equal(t, tc.tag.Color, payload["color"])
				}
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(serverTagData))
			})

			tag, err := tc.write(client, tc.tag)
			if !assert.NoError(t, err) {
				return
			}
			assert.Equal(t, &expectedTag, tag)
		})
	}
}

func TestDeleteTag(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/tags/"+testTagId+"/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		w.WriteHeader(http.StatusNoContent)
	})
	assert.NoError(t, client.DeleteTag(testTagId))
}

func TestTagAPIErrors(t *testing.T) {
	for _, tc := range []struct {
		name   string
		status int
		call   func(*IORiverClient) error
	}{
		{"get missing", http.StatusNotFound, func(c *IORiverClient) error { _, err := c.GetTag(testTagId); return err }},
		{"list forbidden", http.StatusForbidden, func(c *IORiverClient) error { _, err := c.ListTags(); return err }},
		{"create duplicate", http.StatusBadRequest, func(c *IORiverClient) error { _, err := c.CreateTag(expectedTag); return err }},
		{"update invalid", http.StatusBadRequest, func(c *IORiverClient) error { _, err := c.UpdateTag(expectedTag); return err }},
		{"delete missing", http.StatusNotFound, func(c *IORiverClient) error { return c.DeleteTag(testTagId) }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			setup()
			defer teardown()

			mux.HandleFunc("/v1/tags/", func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "tag operation failed", tc.status)
			})

			err := tc.call(client)
			var apiErr *APIError
			if !assert.True(t, errors.As(err, &apiErr), "expected APIError, got %v", err) {
				return
			}
			assert.Equal(t, tc.status, apiErr.StatusCode)
			assert.Contains(t, apiErr.Details, "tag operation failed")
		})
	}
}
