package article

import (
	"CRUD-Operation/conn"
	"errors"
	"net/http"
	"time"

	article "CRUD-Operation/models/article"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// ArticleCollection constant
const ArticleCollection = "article"

var (
	errNotExist        = errors.New("Articles are not existing")
	errInvalidID       = errors.New("Invalid ID")
	errInvalidBody     = errors.New("Invalid request body")
	errInsertionFailed = errors.New("Error in the article insertion")
	errUpdationFailed  = errors.New("Error in the article updation")
	errDeletionFailed  = errors.New("Error in the article deletion")
)

// GetAllArticle Endpoint
func GetAllArticle(c *gin.Context) {
	// Get DB from Mongo Config
	db := conn.GetMongoDB()
	articles := article.Articles{}
	err := db.C(ArticleCollection).Find(bson.M{}).All(&articles)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": errNotExist.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "articles": &articles})
}

// GetArticle Endpoint
func GetArticle(c *gin.Context) {
	// Get Param
	var id = bson.ObjectIdHex(c.Param("id"))
	article, err := article.ArticleInfo(id, ArticleCollection)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": errInvalidID.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "article": &article})
}

// CreateArticle Endpoint
func CreateArticle(c *gin.Context) {
	// Get DB from Mongo Config
	db := conn.GetMongoDB()
	article := article.Article{}
	err := c.Bind(&article)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		//c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
		return
	}
	article.ID = bson.NewObjectId()
	article.CreatedAt = time.Now()
	article.UpdatedAt = time.Now()
	err = db.C(ArticleCollection).Insert(article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": errInsertionFailed.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success", "article": &article})
}

// UpdateArticle Endpoint
func UpdateArticle(c *gin.Context) {
	// Get DB from Mongo Config
	db := conn.GetMongoDB()
	// Get Param
	var id = bson.ObjectIdHex(c.Param("id"))
	existingArticle, err := article.ArticleInfo(id, ArticleCollection)
	if err != nil {
		//c.AbortWithError(http.StatusNotFound, err)
		c.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": errInvalidID.Error()})
		return
	}
	// article := article.Article{}
	err = c.Bind(&existingArticle)
	if err != nil {
		//c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
		return
	}
	existingArticle.ID = id
	existingArticle.UpdatedAt = time.Now()
	err = db.C(ArticleCollection).Update(bson.M{"_id": &id}, existingArticle)
	if err != nil {
		//c.AbortWithError(http.StatusInternalServerError, err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": errUpdationFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "article": &existingArticle})
}

// DeleteArticle Endpoint
func DeleteArticle(c *gin.Context) {
	// Get DB from Mongo Config
	db := conn.GetMongoDB()
	// Get Param
	var id = bson.ObjectIdHex(c.Param("id"))
	err := db.C(ArticleCollection).Remove(bson.M{"_id": &id})
	if err != nil {
		//c.AbortWithError(http.StatusNotFound, err)
		c.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": errDeletionFailed.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"status": "success", "message": "Article deleted successfully"})
}
