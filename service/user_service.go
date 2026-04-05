package service

import (
	"goblog/config"
	"goblog/model"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Create(user *model.User) error {
	return config.DB.Create(user).Error
}

func (s *UserService) GetByID(id uint) (*model.User, error) {
	var user model.User
	err := config.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) GetAll() ([]model.User, error) {
	var users []model.User
	err := config.DB.Find(&users).Error
	return users, err
}

func (s *UserService) Update(user *model.User) error {
	return config.DB.Save(user).Error
}

func (s *UserService) Delete(id uint) error {
	return config.DB.Delete(&model.User{}, id).Error
}
