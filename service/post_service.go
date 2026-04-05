package service

import (
	"goblog/config"
	"goblog/model"
)

type PostService struct{}

func NewPostService() *PostService {
	return &PostService{}
}

func (s *PostService) Create(post *model.Post) error {
	return config.DB.Create(post).Error
}

func (s *PostService) GetByID(id uint) (*model.Post, error) {
	var post model.Post
	err := config.DB.Preload("User").First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (s *PostService) GetAll() ([]model.Post, error) {
	var posts []model.Post
	err := config.DB.Preload("User").Find(&posts).Error
	return posts, err
}

func (s *PostService) Update(post *model.Post) error {
	return config.DB.Save(post).Error
}

func (s *PostService) Delete(id uint) error {
	return config.DB.Delete(&model.Post{}, id).Error
}
