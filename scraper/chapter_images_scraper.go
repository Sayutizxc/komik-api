package scraper

import (
	"strings"

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
		e.ForEach("ol.comment-list > li.comment", func(_ int, element *colly.HTMLElement) {
			var comment model.Comment
			comment.Name = element.ChildText("li.depth-1 > article > div > div.comment-author > h6")
			comment.ProfilePic = element.ChildAttr("li.depth-1 img", "src")
			comment.Comment = element.ChildText("li.depth-1 > article > div > div.comment-content > p")
			comment.Date = element.ChildText("li.depth-1 > article > div > div.comment-metadata > a")
			chapterImages.Comments = append(chapterImages.Comments, comment)
		})
	})

	url = strings.Replace(url, "?style=list", "", -1)
	if err := c.Visit(url + "?style=list"); err != nil {
		return model.ChapterImages{}, err
	}

	return chapterImages, nil
}
