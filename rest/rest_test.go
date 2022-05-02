package rest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocumentHandler(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mux := createHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	var data []urlDescription
	err := json.Unmarshal(res.Body.Bytes(), &data)
	assert.Nil(err)

	url, _ := url("/").MarshalText()
	assert.Equal(url, []byte(data[0].URL))
	assert.Equal("GET", data[0].Method)
	assert.Equal("See Documentation", data[0].Description)
	assert.Equal("", data[0].Payload)
}
