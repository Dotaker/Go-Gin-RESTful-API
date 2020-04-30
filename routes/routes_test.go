package routes_test

import (
	"CRUD-Operation/routes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/* Users Tests */
func TestGetAllUsersRoute(t *testing.T) {
	router := routes.StartGin()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestCreateUserRoute(t *testing.T) {
	router := routes.StartGin()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/users", strings.NewReader(`{ "Name":"admin","Address":"Home"}`))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}

func TestGetUserRoute(t *testing.T) {
	router := routes.StartGin()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/users/5e90a2c3b82d3752a8baae69", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestUpdateUserRoute(t *testing.T) {
	router := routes.StartGin()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/users/5e90a2c3b82d3752a8baae69", strings.NewReader(`{ "Name":"Quentin","Address":"Venelles"}`))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestDeleteUserRoute(t *testing.T) {
	router := routes.StartGin()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/users/5e90a33fb82d37546126e879", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	router.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)
}

/* Articles Tests */
func TestGetAllArticlesRoute(t *testing.T) {
	router := routes.StartGin()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/articles", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestCreateArticleRoute(t *testing.T) {
	router := routes.StartGin()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/articles", strings.NewReader(`{ "Author":"Quentin","Description":"Hi!"}`))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}

func TestGetArticleRoute(t *testing.T) {
	router := routes.StartGin()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/articles/5e90dc1ab82d37a717fb9fd9", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestUpdateArticleRoute(t *testing.T) {
	router := routes.StartGin()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/articles/5e90dc1ab82d37a717fb9fd9", strings.NewReader(`{"Author":"Quentin","Description":"Hello welcome"}`))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestDeleteArticleRoute(t *testing.T) {
	router := routes.StartGin()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/articles/5e90e306b82d37b53456c179", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	router.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)
}
