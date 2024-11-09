package services

import (
	"github.com/Dominux/gotcha/internal/models"
	"github.com/Dominux/gotcha/internal/repositories"
)

type LinkService struct {
	repo *repositories.LinkRepository
}

func NewLinkService(repo *repositories.LinkRepository) *LinkService {
	return &LinkService{repo}
}

func (s *LinkService) Create(linkData *models.LinkDataModel) string {
	return s.repo.Create(linkData)
}

func (s *LinkService) Get(id string) (*models.LinkDataModel, error) {
	linkData, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}

	if linkData.FollowingsLeft == 1 {
		s.repo.Delete(id)
	} else {
		linkData.FollowingsLeft -= 1
	}

	return linkData, nil
}
