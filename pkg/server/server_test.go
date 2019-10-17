package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type response map[string]interface{}

func performRequest(r http.Handler, method, path string, payload []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(payload))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestAddMessageConcurrent(t *testing.T) {
	expectedBody := response{
		"status": "ok",
	}

	responseChan := make(chan *response)

	defer close(responseChan)

	for i := 0; i <= 3; i++ {
		go func(t *testing.T) {
			payload, err := json.Marshal(map[string]interface{}{
				"msg": "Hello",
				"ts":  time.Now().Unix(),
			})

			router := NewRouter(5)

			w := performRequest(router, "POST", "/message", payload)
			var res response
			err = json.Unmarshal([]byte(w.Body.String()), &res)
			assert.Nil(t, err)
			responseChan <- &res
		}(t)

	}

	for i := 0; i <= 3; i++ {
		response := <-responseChan
		assert.Equal(t, expectedBody, *response)
	}
}
