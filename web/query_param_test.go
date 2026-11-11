package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func HandlerSingleQueryParam(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("search")
	fmt.Fprintf(w, "search: %s", q)
}

func TestSingleQueryParam(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8080/dashboard?search=joko", nil)
	rec := httptest.NewRecorder()

	HandlerSingleQueryParam(rec, req)

	res := rec.Result()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	bodyStr := string(body)
	assert.Equal(t, "search: joko", bodyStr, "body is not nil")
}

func HandlerMultipleQueryParam(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	city := r.URL.Query().Get("city")
	fmt.Fprintf(w, "search: %s in %s", search, city)
}

func TestMultipleQueryParam(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8080/dashboard?search=joko&city=solo", nil)
	rec := httptest.NewRecorder()

	HandlerMultipleQueryParam(rec, req)

	res := rec.Result()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	bodyStr := string(body)
	assert.Equal(t, "search: joko in solo", bodyStr, "body is not nil")
}

func HandlerMultipleValuesQueryParam(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	search := q["search"]
	fmt.Fprintf(w, "search: %s", strings.Join(search, " "))
}

func TestMultipleValuesQueryParam(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8080/dashboard?search=joko&search=solo&search=mulyono", nil)
	rec := httptest.NewRecorder()

	HandlerMultipleValuesQueryParam(rec, req)

	res := rec.Result()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	bodyStr := string(body)
	assert.Equal(t, "search: joko solo mulyono", bodyStr, "body is not nil")
}
