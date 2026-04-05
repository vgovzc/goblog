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

var userService = service.NewUserService()

// @Summary Create user
// @Description Create a new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User info"
// @Success 200 {object} model.User
// @Router /api/user/create [post]
type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Code     string `json:"code" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Age      uint   `json:"age"`
}

type UpdateUserRequest struct {
	Username string `json:"username"`
	Code     string `json:"code"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      uint   `json:"age"`
}

func CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	user := &model.User{
		Username: req.Username,
		Code:     req.Code,
		Email:    req.Email,
		Password: req.Password,
		Age:      req.Age,
	}

	if err := userService.Create(user); err != nil {
		util.Error(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	util.Success(c, user)
}

// @Summary Get user by ID
// @Description Get user by ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} model.User
// @Router /api/user/{id} [get]
func GetUserByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.Error(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := userService.GetByID(uint(id))
	if err != nil {
		util.Error(c, http.StatusNotFound, "User not found")
		return
	}

	util.Success(c, user)
}

// @Summary Get user list
// @Description Get all users
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {array} model.User
// @Router /api/user/list [get]
func GetUserList(c *gin.Context) {
	users, err := userService.GetAll()
	if err != nil {
		util.Error(c, http.StatusInternalServerError, "Failed to get users")
		return
	}

	util.Success(c, users)
}

// @Summary Update user
// @Description Update user by ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body UpdateUserRequest true "User info"
// @Success 200 {object} model.User
// @Router /api/user/{id} [put]
func UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.Error(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := userService.GetByID(uint(id))
	if err != nil {
		util.Error(c, http.StatusNotFound, "User not found")
		return
	}

	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Code != "" {
		user.Code = req.Code
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		user.Password = req.Password
	}
	if req.Age > 0 {
		user.Age = req.Age
	}

	if err := userService.Update(user); err != nil {
		util.Error(c, http.StatusInternalServerError, "Failed to update user")
		return
	}

	util.Success(c, user)
}

// @Summary Delete user
// @Description Delete user by ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200
// @Router /api/user/{id} [delete]
func DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.Error(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	if err := userService.Delete(uint(id)); err != nil {
		util.Error(c, http.StatusInternalServerError, "Failed to delete user")
		return
	}

	util.Success(c, nil)
}
