package service

import (
	"github.com/sayutizxc/klikmanga-scraper/model"
	"github.com/sayutizxc/klikmanga-scraper/repository"
)

func ListComicService(page string, limit string) ([]model.Comic, error) {

	comics, err := repository.ListComicRepository(page, limit)
	if err != nil {
		return nil, err
	}

	return comics, nil
}
