package routes_test

import (
	user "CRUD-Operation/models/user"
	"CRUD-Operation/routes"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	user := user.User{}
	u, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(u))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}
