package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const COOKIE_NAME = "x-jkw-cookie"
const COOKIE_VALUE = "jokowiii"
const COOKIE_PATH = "/"

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)

	cookie.Name = COOKIE_NAME
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = COOKIE_PATH

	http.SetCookie(w, cookie)
	fmt.Fprintf(w, "Success!")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(COOKIE_NAME)
	if err != nil {
		fmt.Fprintf(w, "Error while getting cookie %s", err)
	} else {
		fmt.Fprintf(w, "Here is the cookie: %s", cookie.Value)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

func TestSetCookie(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/set-cookie?name=jokowiii", nil)
	rec := httptest.NewRecorder()

	SetCookie(rec, req)

	res := rec.Result().Cookies()

	assert.Equal(t, COOKIE_NAME, res[0].Name, "cookie key is not equal!")
	assert.Equal(t, COOKIE_VALUE, res[0].Value, "cookie value is not equal!")
}

func TestGetCookie(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/get-cookie", nil)
	cookie := new(http.Cookie)
	cookie.Name = COOKIE_NAME
	cookie.Value = COOKIE_VALUE
	cookie.Path = COOKIE_PATH
	req.AddCookie(cookie)

	rec := httptest.NewRecorder()

	GetCookie(rec, req)

	res, _ := io.ReadAll(rec.Result().Body)
	fmt.Println("res", string(res))

	assert.Equal(t, "Here is the cookie: "+COOKIE_VALUE, string(res), "cookie value is not equal!")
}
