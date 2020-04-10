package model

import (
	"CRUD-Operation/conn"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// User structure
type User struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string        `bson:"name"`
	Address   string        `bson:"address"`
	Age       int           `bson:"age"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`
}

// Users list
type Users []User

// UserInfo model function
func UserInfo(id bson.ObjectId, userCollection string) (User, error) {
	db := conn.GetMongoDB()
	user := User{}
	err := db.C(userCollection).Find(bson.M{"_id": &id}).One(&user)
	return user, err
}
