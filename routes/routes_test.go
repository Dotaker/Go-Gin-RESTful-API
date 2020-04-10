package routes_test

import (
	"CRUD-Operation/routes"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestCreateUserRoute(t *testing.T) {
	router := routes.StartGin()

	w := httptest.NewRecorder()
	reqBody, _ := json.Marshal(map[string]string{"Name": "Admin", "Address": "Home"})
	req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}
