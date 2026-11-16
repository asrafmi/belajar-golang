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

func SimpleFormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() // or directly using r.PostFormValue()
	if err != nil {
		panic(err)
	}

	firstName := r.PostForm.Get("first_name")
	lastName := r.PostForm.Get("last_name")

	fmt.Fprintf(w, "Hi %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	formBody := strings.NewReader("first_name=Joko&last_name=Widodo")
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", formBody)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	rec := httptest.NewRecorder()

	SimpleFormPost(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	assert.Equal(t, "Hi Joko Widodo", string(body), "Form Body is not equal")
}
