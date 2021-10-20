package scraper

import (
	"github.com/gocolly/colly"
	"github.com/sayutizxc/klikmanga-scraper/model"
)

func ChapterImagesScraper(url string) (model.ChapterImages, error) {
	chapterImages, err := getChapterImages(url)
	if err != nil {
		return model.ChapterImages{}, err
	}
	return chapterImages, nil
}

func getChapterImages(url string) (model.ChapterImages, error) {
	var chapterImages model.ChapterImages
	chapterImages.Url = url
	c := colly.NewCollector()

	c.OnHTML("div.main-col-inner", func(e *colly.HTMLElement) {
		chapterImages.Chapter = e.ChildText("div > div.c-breadcrumb > ol > li.active")
		chapterImages.Prev = e.ChildAttr("div.nav-previous > a", "href")
		chapterImages.Next = e.ChildAttr("div.nav-next > a", "href")
		e.ForEach(".page-break.no-gaps", func(i int, element *colly.HTMLElement) {
			chapterImages.Images = append(chapterImages.Images, element.ChildAttr("img", "src"))
		})
		e.ForEach("li.comment", func(_ int, element *colly.HTMLElement) {
			var comment model.Comment
			comment.Name = element.ChildText("div.comment-author > h6")
			comment.ProfilePic = element.ChildAttr("div.comment-avatar > img", "src")
			comment.Comment = element.ChildText("div.comment-content > p")
			comment.Date = element.ChildText("div.comment-metadata > a")
			chapterImages.Comments = append(chapterImages.Comments, comment)
		})
	})

	if err := c.Visit(url + "?style=list"); err != nil {
		return model.ChapterImages{}, err
	}

	return chapterImages, nil
}
