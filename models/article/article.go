package model

import (
	"CRUD-Operation/conn"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Article structure
type Article struct {
	ID          bson.ObjectId `bson:"_id"`
	Author      string        `bson:"author"`
	Description string        `bson:"description"`
	CreatedAt   time.Time     `bson:"created_at"`
	UpdatedAt   time.Time     `bson:"updated_at"`
}

// Articles list
type Articles []Article

// ArticleInfo model function
func ArticleInfo(id bson.ObjectId, articleCollection string) (Article, error) {
	db := conn.GetMongoDB()
	article := Article{}
	err := db.C(articleCollection).Find(bson.M{"_id": &id}).One(&article)
	return article, err
}
