package ioriver

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCallApi_BackgroundTaskCompleted(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/test_async_completed", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Header().Set("x-request-id", "req-complete")
		w.Header().Set("x-background-task-id", "42")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"ok":true}`)
	})

	mux.HandleFunc("/v1/async_task_by_id/42/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"id":42,"status":"COMPLETED","progress":100}`)
	})

	resp, err := client.CallApi("test_async_completed", http.MethodGet, CallParams{})
	if assert.NoError(t, err) {
		assert.NotNil(t, resp)
		resp.Body.Close()
	}
}

func TestCallApi_BackgroundTaskError(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/test_async_error", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Header().Set("x-request-id", "req-error")
		w.Header().Set("x-background-task-id", "43")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"ok":true}`)
	})

	mux.HandleFunc("/v1/async_task_by_id/43/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"id":43,"status":"ERROR","message":"boom","details":"stack"}`)
	})

	resp, err := client.CallApi("test_async_error", http.MethodGet, CallParams{})
	assert.Nil(t, resp)
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "request failed: boom")
		assert.Contains(t, err.Error(), "request-id: req-error")
		assert.Contains(t, err.Error(), "details: stack")
	}
}

func TestCallApi_BackgroundTaskNotFound(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/test_async_not_found", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Header().Set("x-request-id", "req-missing")
		w.Header().Set("x-background-task-id", "44")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"ok":true}`)
	})

	mux.HandleFunc("/v1/async_task_by_id/44/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(http.StatusNotFound)
	})

	resp, err := client.CallApi("test_async_not_found", http.MethodGet, CallParams{})
	assert.Nil(t, resp)
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "async task not found: id=44")
		assert.Contains(t, err.Error(), "request-id=req-missing")
	}
}

func TestCallApi_BackgroundTaskRequestFailure(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/test_async_failure", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Header().Set("x-request-id", "req-fail")
		w.Header().Set("x-background-task-id", "45")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"ok":true}`)
	})

	mux.HandleFunc("/v1/async_task_by_id/45/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Header().Set("x-request-id", "async-req-fail")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"error":"db down"}`)
	})

	resp, err := client.CallApi("test_async_failure", http.MethodGet, CallParams{})
	assert.Nil(t, resp)
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "500 Internal Server Error")
		assert.Contains(t, err.Error(), "request-id: async-req-fail")
		assert.Contains(t, err.Error(), "db down")
	}
}
