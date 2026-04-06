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

var commentService = service.NewCommentService()

type CreateCommentRequest struct {
	Content string `json:"content" binding:"required"`
	PostID  uint   `json:"post_id" binding:"required"`
	UserID  uint   `json:"user_id" binding:"required"`
}

type UpdateCommentRequest struct {
	Content string `json:"content"`
}

// @Summary Create comment
// @Description Create a new comment
// @Tags comment
// @Accept json
// @Produce json
// @Param comment body CreateCommentRequest true "Comment info"
// @Success 200 {object} model.Comment
// @Router /api/comment/create [post]
func CreateComment(c *gin.Context) {
	var req CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	comment := &model.Comment{
		Content: req.Content,
		PostID:  req.PostID,
		UserID:  req.UserID,
	}

	if err := commentService.Create(comment); err != nil {
		util.Error(c, http.StatusInternalServerError, "Failed to create comment")
		return
	}

	util.Success(c, comment)
}

// @Summary Get comment by ID
// @Description Get comment by ID
// @Tags comment
// @Accept json
// @Produce json
// @Param id path int true "Comment ID"
// @Success 200 {object} model.Comment
// @Router /api/comment/{id} [get]
func GetCommentByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.Error(c, http.StatusBadRequest, "Invalid comment ID")
		return
	}

	comment, err := commentService.GetByID(uint(id))
	if err != nil {
		util.Error(c, http.StatusNotFound, "Comment not found")
		return
	}

	util.Success(c, comment)
}

// @Summary Get comment list
// @Description Get all comments
// @Tags comment
// @Accept json
// @Produce json
// @Success 200 {array} model.Comment
// @Router /api/comment/list [get]
func GetCommentList(c *gin.Context) {
	comments, err := commentService.GetAll()
	if err != nil {
		util.Error(c, http.StatusInternalServerError, "Failed to get comments")
		return
	}

	util.Success(c, comments)
}

// @Summary Update comment
// @Description Update comment by ID
// @Tags comment
// @Accept json
// @Produce json
// @Param id path int true "Comment ID"
// @Param comment body UpdateCommentRequest true "Comment info"
// @Success 200 {object} model.Comment
// @Router /api/comment/{id} [put]
func UpdateComment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.Error(c, http.StatusBadRequest, "Invalid comment ID")
		return
	}

	var req UpdateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	comment, err := commentService.GetByID(uint(id))
	if err != nil {
		util.Error(c, http.StatusNotFound, "Comment not found")
		return
	}

	if req.Content != "" {
		comment.Content = req.Content
	}

	if err := commentService.Update(comment); err != nil {
		util.Error(c, http.StatusInternalServerError, "Failed to update comment")
		return
	}

	util.Success(c, comment)
}

// @Summary Delete comment
// @Description Delete comment by ID
// @Tags comment
// @Accept json
// @Produce json
// @Param id path int true "Comment ID"
// @Success 200
// @Router /api/comment/{id} [delete]
func DeleteComment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.Error(c, http.StatusBadRequest, "Invalid comment ID")
		return
	}

	if err := commentService.Delete(uint(id)); err != nil {
		util.Error(c, http.StatusInternalServerError, "Failed to delete comment")
		return
	}

	util.Success(c, nil)
}
