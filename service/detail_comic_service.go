package service

import (
	"github.com/sayutizxc/klikmanga-scraper/model"
	"github.com/sayutizxc/klikmanga-scraper/repository"
)

func DetailComicService(url string) (model.DetailComic, error) {
	detailComic, err := repository.DetailComicRepository(url)
	if err != nil {
		return model.DetailComic{}, err
	}
	return detailComic, nil
}
