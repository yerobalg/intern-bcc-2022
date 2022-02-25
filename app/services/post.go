package services

import (
	"clean-arch-2/app/repositories"
	"clean-arch-2/app/models"
)

type PostService struct {
	repo repositories.PostRepository
}

func NewPostService(postRepo repositories.PostRepository) PostService {
	return PostService{repo: postRepo}
}

func (s PostService) Save(model *models.Post)  (error) {
	return s.repo.Save(model)
}