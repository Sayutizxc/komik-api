package service

import (
	"github.com/sayutizxc/klikmanga-scraper/model"
	"github.com/sayutizxc/klikmanga-scraper/repository"
)

func SearchComicService(s string, page string, limit string) ([]model.Comic, error) {
	comics, err := repository.SearchComicRepository(s, page, limit)
	if err != nil {
		return nil, err
	}
	return comics, nil
}
