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
	return s.repo.Get(id)
}
