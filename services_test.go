package ioriver

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
