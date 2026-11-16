package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SimpleHeader(w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("content-type")
	w.Header().Add("content-type", header+"-header")
	fmt.Fprint(w, header+"-body")
}

func TestSimpleHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	req.Header.Add("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	SimpleHeader(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)
	fmt.Println("body", string(body))
	header := res.Header.Get("content-type")
	fmt.Println("header", header)

	assert.Equal(t, "application/json-header", header, "content-type from body is not equal")
	assert.Equal(t, "application/json-body", string(body), "content-type from header is not equal")
}
