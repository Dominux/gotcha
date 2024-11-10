package services

import (
	"fmt"
	"time"

	"github.com/Dominux/gotcha/internal/models"
	"github.com/Dominux/gotcha/internal/repositories"
)

const Day = time.Hour * 24

type LinkService struct {
	repo    *repositories.LinkRepository
	urlBase *string
}

func NewLinkService(repo *repositories.LinkRepository, urlBase *string) *LinkService {
	return &LinkService{repo, urlBase}
}

func (s *LinkService) Create(linkData *models.LinkDataModel) string {
	id := s.repo.Create(linkData)
	return fmt.Sprintf("%s/l/%s", *s.urlBase, id)
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

func (s *LinkService) RunLinksRemovingCycle() {
	println("Ran links removing cycle")

	for {
		time.Sleep(Day)

		var deletedCounter uint
		f := func(id string, linkData *models.LinkDataModel) {
			if linkData.DaysLeft == 1 {
				s.repo.Delete(id)
				deletedCounter += 1
			} else {
				linkData.DaysLeft -= 1
			}
		}

		s.repo.Map(f)

		fmt.Printf("Removed %d outdated links\n", deletedCounter)
	}
}
