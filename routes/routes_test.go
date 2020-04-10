package routes_test

import (
	"CRUD-Operation/routes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	user "CRUD-Operation/models/user"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string        `bson:"name"`
	Address   string        `bson:"address"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`
}

func TestGetAllUsersRoute(t *testing.T) {
	router := routes.StartGin()
	users := user.Users{}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, gin.H{"status": "success", "users": &users}, gin.H{"status": "success", "users": &users})
}

func TestCreateUserRoute(t *testing.T) {
	router := routes.StartGin()
	user := user.User{}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/users", strings.NewReader(`{ "Name":"admin","Address":"Home"}`))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, gin.H{"status": "success", "user": &user}, gin.H{"status": "success", "user": &user})
}

func TestGetUserRoute(t *testing.T) {
	router := routes.StartGin()
	user := user.User{}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/users/5e90a2c3b82d3752a8baae69", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, gin.H{"status": "success", "user": &user}, gin.H{"status": "success", "user": &user})
}

func TestUpdateUserRoute(t *testing.T) {
	router := routes.StartGin()
	user := user.User{}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/users/5e90a2c3b82d3752a8baae69", strings.NewReader(`{ "Name":"Quentin","Address":"Venelles"}`))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, gin.H{"status": "success", "user": &user}, gin.H{"status": "success", "user": &user})
}

func TestDeleteUserRoute(t *testing.T) {
	router := routes.StartGin()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/users/5e90c641b82d377a1e08b291", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, gin.H{"status": "success", "message": "User deleted successfully"}, gin.H{"status": "success", "message": "User deleted successfully"})
}
