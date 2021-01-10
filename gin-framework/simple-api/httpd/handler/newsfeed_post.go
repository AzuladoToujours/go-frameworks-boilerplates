package handler

import (
	"net/http"

	"github.com/azuladotoujours/go-framework-tests/gin-framework/simple-api/platform/newsfeed"
	"github.com/gin-gonic/gin"
)

type newsfeedPostRequest struct {
	Title string `json: "title"`
	Post  string `json: "post"`
}

func NewsfeedPost(feed newsfeed.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := newsfeedPostRequest{}
		c.Bind(&requestBody)
		item := newsfeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Title,
		}
		feed.Add(item)
		c.Status(http.StatusNoContent)
	}
}
