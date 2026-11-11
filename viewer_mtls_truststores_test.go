package ioriver

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestViewerMTLSTrustStoreCRUD(t *testing.T) {
	expected := ViewerMTLSTrustStore{
		Id:      testObjectId,
		Account: "cdddb6ea-b740-4828-aa41-c49ed7dfab42",
		Name:    "client-ca",
		Bundle:  "-----BEGIN CERTIFICATE-----\nclient-ca\n-----END CERTIFICATE-----\n",
	}
	for _, operation := range []string{"list", "get", "create", "delete"} {
		t.Run(operation, func(t *testing.T) {
			method, path := http.MethodGet, "/v1/viewer-mtls-truststore/"
			switch operation {
			case "get":
				path += testObjectId + "/"
			case "create":
				method = http.MethodPost
			case "delete":
				method = http.MethodDelete
				path += testObjectId + "/"
			}
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != method || r.URL.Path != path {
					t.Errorf("request = %s %s, want %s %s", r.Method, r.URL.Path, method, path)
					http.Error(w, "unexpected request", http.StatusBadRequest)
					return
				}
				if operation == "create" {
					var body map[string]interface{}
					if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
						t.Errorf("decode request: %v", err)
						http.Error(w, "invalid JSON", http.StatusBadRequest)
						return
					}
					want := map[string]interface{}{"name": expected.Name, "bundle": expected.Bundle}
					if !reflect.DeepEqual(body, want) {
						t.Errorf("create body = %#v, want %#v", body, want)
					}
					w.WriteHeader(http.StatusCreated)
				}
				switch operation {
				case "list":
					_ = json.NewEncoder(w).Encode([]ViewerMTLSTrustStore{expected})
				case "delete":
					w.WriteHeader(http.StatusNoContent)
				default:
					_ = json.NewEncoder(w).Encode(expected)
				}
			}))
			defer server.Close()
			client := NewClient("test-token")
			client.EndpointUrl = server.URL + "/"
			switch operation {
			case "list":
				got, err := client.ListViewerMTLSTrustStores()
				if err != nil || !reflect.DeepEqual(got, []ViewerMTLSTrustStore{expected}) {
					t.Fatalf("list = %#v, %v", got, err)
				}
			case "get", "create":
				var got *ViewerMTLSTrustStore
				var err error
				if operation == "get" {
					got, err = client.GetViewerMTLSTrustStore(testObjectId)
				} else {
					got, err = client.CreateViewerMTLSTrustStore(ViewerMTLSTrustStore{
						Name: expected.Name, Bundle: expected.Bundle,
					})
				}
				if err != nil || got == nil || *got != expected {
					t.Fatalf("%s = %#v, %v", operation, got, err)
				}
			case "delete":
				if err := client.DeleteViewerMTLSTrustStore(testObjectId); err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}

func TestViewerMTLSTrustStoreErrors(t *testing.T) {
	for _, status := range []int{http.StatusBadRequest, http.StatusNotFound, http.StatusForbidden} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "trust store is in use", status)
			}))
			defer server.Close()
			client := NewClient("test-token")
			client.EndpointUrl = server.URL + "/"
			operations := []func() error{
				func() error { _, err := client.ListViewerMTLSTrustStores(); return err },
				func() error { _, err := client.GetViewerMTLSTrustStore(testObjectId); return err },
				func() error {
					_, err := client.CreateViewerMTLSTrustStore(ViewerMTLSTrustStore{Name: "test", Bundle: "bundle"})
					return err
				},
				func() error { return client.DeleteViewerMTLSTrustStore(testObjectId) },
			}
			for i, operation := range operations {
				var apiErr *APIError
				if err := operation(); !errors.As(err, &apiErr) || apiErr.StatusCode != status {
					t.Errorf("operation %d: expected API error %d, got %v", i, status, err)
				}
			}
		})
	}
}
