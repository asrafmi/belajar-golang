package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Name is required")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hi %s", name)
	}
}

func TestResponseCodeValid(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=Joko", nil)
	rec := httptest.NewRecorder()

	ResponseCode(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	assert.Equal(t, "Hi Joko", string(body), "body is not equal")
	assert.Equal(t, http.StatusOK, res.StatusCode, "status code is not equal")
}

func TestResponseCodeInvalid(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	ResponseCode(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	assert.Equal(t, "Name is required", string(body), "body is not equal")
	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "status code is not equal")
}
