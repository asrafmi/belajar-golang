package golang_web

import (
	"embed"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SimpleHTML(w http.ResponseWriter, r *http.Request) {
	html := "<html><body>{{.}}</body></html>"
	template, err := template.New("SIMPLE").Parse(html)
	if err != nil {
		panic(err)
	}

	template.ExecuteTemplate(w, "SIMPLE", "Hello World!")
}

func TestSimpleHTML(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	res := httptest.NewRecorder()

	SimpleHTML(res, req)

	body, _ := io.ReadAll(res.Result().Body)
	assert.Equal(t, "<html><body>Hello World!</body></html>", string(body), "template is not equal")
}

func SimpleHTMLFile(w http.ResponseWriter, r *http.Request) {

	template := template.Must(template.ParseFiles("./templates/index.gohtml"))

	template.ExecuteTemplate(w, "index.gohtml", "Hello World!")
}

func TestSimpleHTMLFile(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	res := httptest.NewRecorder()

	SimpleHTMLFile(res, req)

	body, _ := io.ReadAll(res.Result().Body)
	assert.Equal(t, "<!DOCTYPE html><html lang='en'><head><meta charset='UTF-8'><meta name='viewport' content='width=device-width, initial-scale=1.0'><title>Hello World!</title></head><body><h1>Hello World!</h1></body></html>", string(body), "template is not equal")
}

func SimpleHTMLDirectory(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseGlob("./templates/*.gohtml"))

	template.ExecuteTemplate(w, "index.gohtml", "Hello World!")
}

func TestSimpleHTMLFileDirectory(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	res := httptest.NewRecorder()

	SimpleHTMLDirectory(res, req)

	body, _ := io.ReadAll(res.Result().Body)
	assert.Equal(t, "<!DOCTYPE html><html lang='en'><head><meta charset='UTF-8'><meta name='viewport' content='width=device-width, initial-scale=1.0'><title>Hello World!</title></head><body><h1>Hello World!</h1></body></html>", string(body), "template is not equal")
}

//go:embed templates/*.gohtml
var templates embed.FS

func SimpleHTMLEmbed(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFS(templates, "templates/*.gohtml"))

	template.ExecuteTemplate(w, "index.gohtml", "Hello World!")
}

func TestSimpleHTMLFileEmbed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	res := httptest.NewRecorder()

	SimpleHTMLEmbed(res, req)

	body, _ := io.ReadAll(res.Result().Body)
	assert.Equal(t, "<!DOCTYPE html><html lang='en'><head><meta charset='UTF-8'><meta name='viewport' content='width=device-width, initial-scale=1.0'><title>Hello World!</title></head><body><h1>Hello World!</h1></body></html>", string(body), "template is not equal")
}
