package repositories

import (
	"errors"

	"github.com/Dominux/gotcha/internal/models"
	"github.com/google/uuid"
)

type LinkRepository struct {
	links map[string]*models.LinkDataModel
}

func NewLinkRepository() *LinkRepository {
	links := make(map[string]*models.LinkDataModel)
	return &LinkRepository{links}
}

func (r *LinkRepository) Create(linkData *models.LinkDataModel) string {
	id := uuid.New().String()
	r.links[id] = linkData
	return id
}

func (r *LinkRepository) Get(id string) (*models.LinkDataModel, error) {
	if link, exists := r.links[id]; exists {
		return link, nil
	} else {
		return nil, errors.New("does not exist")
	}
}
