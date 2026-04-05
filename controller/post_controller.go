// Package controller provides HTTP handlers for the blog API
// @Title Blog API
// @Description Blog backend API with user, post, and comment management
// @Version 1.0
package controller

import (
	"net/http"
	"strconv"

	"goblog/model"
	"goblog/service"
	"goblog/util"

	"github.com/gin-gonic/gin"
)

var postService = service.NewPostService()

// @Summary Create article
// @Description Create a new article
// @Tags article
// @Accept json
// @Produce json
// @Param article body CreatePostRequest true "Article info"
// @Success 200 {object} model.Post
// @Router /api/article/create [post]
type CreatePostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id" binding:"required"`
}

type UpdatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func CreateArticle(c *gin.Context) {
	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	post := &model.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  req.UserID,
	}

	if err := postService.Create(post); err != nil {
		util.Error(c, http.StatusInternalServerError, "Failed to create post")
		return
	}

	util.Success(c, post)
}

// @Summary Get article by ID
// @Description Get article by ID
// @Tags article
// @Accept json
// @Produce json
// @Param id path int true "Article ID"
// @Success 200 {object} model.Post
// @Router /api/article/{id} [get]
func GetArticleByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.Error(c, http.StatusBadRequest, "Invalid article ID")
		return
	}

	post, err := postService.GetByID(uint(id))
	if err != nil {
		util.Error(c, http.StatusNotFound, "Article not found")
		return
	}

	util.Success(c, post)
}

// @Summary Get article list
// @Description Get all articles
// @Tags article
// @Accept json
// @Produce json
// @Success 200 {array} model.Post
// @Router /api/article/list [get]
func GetArticleList(c *gin.Context) {
	posts, err := postService.GetAll()
	if err != nil {
		util.Error(c, http.StatusInternalServerError, "Failed to get articles")
		return
	}

	util.Success(c, posts)
}

// @Summary Update article
// @Description Update article by ID
// @Tags article
// @Accept json
// @Produce json
// @Param id path int true "Article ID"
// @Param article body UpdatePostRequest true "Article info"
// @Success 200 {object} model.Post
// @Router /api/article/{id} [put]
func UpdateArticle(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.Error(c, http.StatusBadRequest, "Invalid article ID")
		return
	}

	var req UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	post, err := postService.GetByID(uint(id))
	if err != nil {
		util.Error(c, http.StatusNotFound, "Article not found")
		return
	}

	if req.Title != "" {
		post.Title = req.Title
	}
	if req.Content != "" {
		post.Content = req.Content
	}

	if err := postService.Update(post); err != nil {
		util.Error(c, http.StatusInternalServerError, "Failed to update article")
		return
	}

	util.Success(c, post)
}

// @Summary Delete article
// @Description Delete article by ID
// @Tags article
// @Accept json
// @Produce json
// @Param id path int true "Article ID"
// @Success 200
// @Router /api/article/{id} [delete]
func DeleteArticle(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.Error(c, http.StatusBadRequest, "Invalid article ID")
		return
	}

	if err := postService.Delete(uint(id)); err != nil {
		util.Error(c, http.StatusInternalServerError, "Failed to delete article")
		return
	}

	util.Success(c, nil)
}
