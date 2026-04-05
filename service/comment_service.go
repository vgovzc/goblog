package service

import (
	"goblog/config"
	"goblog/model"
)

type CommentService struct{}

func NewCommentService() *CommentService {
	return &CommentService{}
}

func (s *CommentService) Create(comment *model.Comment) error {
	return config.DB.Create(comment).Error
}

func (s *CommentService) GetByID(id uint) (*model.Comment, error) {
	var comment model.Comment
	err := config.DB.Preload("User").Preload("Post").First(&comment, id).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (s *CommentService) GetAll() ([]model.Comment, error) {
	var comments []model.Comment
	err := config.DB.Preload("User").Preload("Post").Find(&comments).Error
	return comments, err
}

func (s *CommentService) Update(comment *model.Comment) error {
	return config.DB.Save(comment).Error
}

func (s *CommentService) Delete(id uint) error {
	return config.DB.Delete(&model.Comment{}, id).Error
}
