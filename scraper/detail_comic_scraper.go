package scraper

import (
	"github.com/gocolly/colly"
	"github.com/sayutizxc/klikmanga-scraper/model"
	"strings"
)

func DetailComicScraper(url string) (model.DetailComic, error) {
	detailComic, err := getDetailKomik(url)
	if err != nil {
		return model.DetailComic{}, err
	}
	return detailComic, nil
}

func getDetailKomik(url string) (model.DetailComic, error) {

	var detailComic model.DetailComic

	c := colly.NewCollector()

	c.OnHTML(".site-content", func(e *colly.HTMLElement) {
		detailComic.Url = url
		detailComic.Title = e.ChildText("div.post-title > h1")
		detailComic.Thumbnail = e.ChildAttr("div.summary_image > a > img", "src")
		detailComic.Rating = e.ChildText("div.post-total-rating > span")
		detailComic.Authors = e.ChildText("div.summary-content > div.author-content > a")
		detailComic.Artists = e.ChildText("div.summary-content > div.artist-content > a")
		e.ForEach("div.summary-content > div.genres-content > a", func(_ int, element *colly.HTMLElement) {
			detailComic.Genres = append(detailComic.Genres, element.Text)
		})
		detailComic.Type = e.ChildText("div.post-content > div:nth-child(9) > div.summary-content")
		detailComic.Release = e.ChildText("div.post-status > div:nth-child(1) > div.summary-content > a")
		detailComic.Status = e.ChildText("div.post-status > div:nth-child(2) > div.summary-content")
		detailComic.Synopsis = e.ChildText("div.c-page-content.style-1 > div > div > div > div.main-col.col-md-8.col-sm-8 > div > div.c-page > div > div.description-summary > div > p")
		e.ForEach("li.wp-manga-chapter", func(_ int, element *colly.HTMLElement) {
			var chapter model.Chapter
			chapter.ChapterUrl = element.ChildAttr("a", "href")
			chapter.Chapter = element.ChildText("a")
			chapter.Date = element.ChildText("span.chapter-release-date")
			if chapter.Date == "" {
				chapter.Date = element.ChildAttr("span.chapter-release-date > span > a", "title")
				chapter.Date = strings.Replace(chapter.Date, "ago", "yang lalu", 1)
			}
			detailComic.Chapters = append(detailComic.Chapters, chapter)
		})
		e.ForEach("li.comment", func(_ int, element *colly.HTMLElement) {
			var comment model.Comment
			comment.Name = element.ChildText("div.comment-author > h6")
			comment.ProfilePic = element.ChildAttr("div.comment-avatar > img", "src")
			comment.Comment = element.ChildText("div.comment-content > p")
			comment.Date = element.ChildText("div.comment-metadata > a")
			detailComic.Comments = append(detailComic.Comments, comment)
		})
	})

	err := c.Visit(url)
	if err != nil {
		return model.DetailComic{}, err
	}
	return detailComic, nil

}
