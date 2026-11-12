package ioriver

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateAttachedFile(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/attached-files/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "token dummyapitoken", r.Header.Get("Authorization"))
		var payload map[string]string
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, map[string]string{"file_type": "HTML", "contents": "<html>blocked</html>"}, payload)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"id":"file-id","account":"account-id","file_type":"HTML","contents":"<html>blocked</html>"}`))
	})
	file, err := client.CreateAttachedFile("HTML", "<html>blocked</html>")
	require.NoError(t, err)
	assert.Equal(t, "file-id", file.ID)
	assert.Equal(t, "<html>blocked</html>", file.Contents)
}

func TestGetAttachedFile(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/attached-files/file-id/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		_, _ = w.Write([]byte(`{"id":"file-id","file_type":"HTML","contents":"blocked"}`))
	})
	file, err := client.GetAttachedFile("file-id")
	require.NoError(t, err)
	assert.Equal(t, "file-id", file.ID)
	assert.Equal(t, "blocked", file.Contents)
}

func TestDeleteAttachedFile(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/attached-files/file-id/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		w.WriteHeader(http.StatusNoContent)
	})
	require.NoError(t, client.DeleteAttachedFile("file-id"))
}

func TestAttachedFileAPIErrors(t *testing.T) {
	for _, tc := range []struct {
		name string
		call func() error
	}{
		{"create", func() error { _, err := client.CreateAttachedFile("HTML", "body"); return err }},
		{"read", func() error { _, err := client.GetAttachedFile("file-id"); return err }},
		{"delete", func() error { return client.DeleteAttachedFile("file-id") }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			setup()
			defer teardown()
			mux.HandleFunc("/v1/attached-files/", func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "attachment is referenced", http.StatusBadRequest)
			})
			err := tc.call()
			var apiErr *APIError
			require.True(t, errors.As(err, &apiErr), "expected APIError, got %v", err)
			assert.Equal(t, http.StatusBadRequest, apiErr.StatusCode)
		})
	}
}
