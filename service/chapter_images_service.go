package service

import (
	"github.com/sayutizxc/klikmanga-scraper/model"
	"github.com/sayutizxc/klikmanga-scraper/repository"
)

func ChapterImagesService(url string) (model.ChapterImages, error) {
	chapterImages, err := repository.ChapterImagesRepository(url)
	if err != nil {
		return model.ChapterImages{}, err
	}
	return chapterImages, nil
}
