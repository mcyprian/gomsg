package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string, payload []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(payload))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestAddMessage(t *testing.T) {
	expectedBody := map[string]interface{}{
		"status": "ok",
	}
	payload, err := json.Marshal(map[string]interface{}{
		"msg": "Hello",
		"ts":  123,
	})

	router := NewRouter(5)

	w := performRequest(router, "POST", "/message", payload)

	var response map[string]interface{}
	err = json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)
	assert.Equal(t, response, expectedBody)
}
