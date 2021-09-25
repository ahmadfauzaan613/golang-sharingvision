package handler

import (
	"golang-sharingvision/article"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handlerPost struct {
	articleService article.Service
}

func PostHandler(articleService article.Service) *handlerPost {
	return &handlerPost{articleService}
}

// Menambahkan Article
func (h *handlerPost) NewPostHandler(c *gin.Context) {
	var articleInput article.ArticleRequest
	err := c.ShouldBindJSON(&articleInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	article, err := h.articleService.NewArticle(articleInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": article,
	})
}

// GET SEMUA DATA
func (h *handlerPost) GetAll(c *gin.Context) {
	articles, err := h.articleService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var responsePost []article.PostsResponse

	for _, r := range articles {
		responsePosts := article.PostsResponse{
			Title:    r.Title,
			Content:  r.Content,
			Category: r.Category,
			Status:   r.Status,
		}
		responsePost = append(responsePost, responsePosts)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": responsePost,
	})
}

// GET ID DATA
func (h *handlerPost) GetID(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	foundarticle, err := h.articleService.FindByID(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	responsePost := article.PostsResponse{
		Title:    foundarticle.Title,
		Content:  foundarticle.Content,
		Category: foundarticle.Category,
		Status:   foundarticle.Status,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": responsePost,
	})
}

func (h *handlerPost) Deleteid(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	deleteArticle, err := h.articleService.ArticleDelete(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": deleteArticle,
	})
}

func (h *handlerPost) UpdatePostHandler(c *gin.Context) {
	var articleInput article.ArticleRequest
	err := c.ShouldBindJSON(&articleInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	article, err := h.articleService.UpdateArticle(id, articleInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": article,
	})
}
