package mogwai

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	type testJSONData struct {
		ID int `json:"id"`
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		data := testJSONData{
			ID: 1,
		}
		Respond(data, w)
	}

	req := httptest.NewRequest("GET", "/some/path", nil)
	res := httptest.NewRecorder()

	handler(res, req)

	assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
	expected, err := json.Marshal(testJSONData{ID: 1})
	assert.NoError(t, err)
	assert.Equal(t, expected, res.Body.Bytes())
}
