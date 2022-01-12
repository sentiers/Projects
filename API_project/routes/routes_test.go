package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	r := Routers()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hello", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "hello", w.Body.String())
}
