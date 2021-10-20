package scraper

import (
	"github.com/gocolly/colly"
	"github.com/sayutizxc/klikmanga-scraper/model"
	"strings"
)

func ListComicScraper(page string, limit string) ([]model.Comic, error) {
	comics, err := getListComic(page, limit)
	if err != nil {
		return nil, err
	}
	return comics, nil
}

func getListComic(page string, limit string) ([]model.Comic, error) {
	// Membuat penampung dari semua komik yang akan diambil
	var comics []model.Comic

	// Membuat collector dari package go-colly
	c := colly.NewCollector()

	c.OnHTML(".page-item-detail.manga", func(e *colly.HTMLElement) {
		var comic model.Comic

		comic.Url = e.ChildAttr("h3 > a", "href")
		comic.Title = e.ChildText("h3 > a")
		comic.Thumbnail = e.ChildAttr("img", "src")
		comic.Score = e.ChildText("span.score")

		e.ForEach(".chapter-item", func(_ int, element *colly.HTMLElement) {
			var chapter model.Chapter
			chapter.ChapterUrl = element.ChildAttr("span.chapter.font-meta > a", "href")
			chapter.Chapter = element.ChildText("span.chapter.font-meta > a")
			chapter.Date = element.ChildText("span.post-on.font-meta")
			if chapter.Date == "" {
				chapter.Date = element.ChildAttr("span.post-on.font-meta > span > a", "title")
				chapter.Date = strings.Replace(chapter.Date, "ago", "yang lalu", 1)
			}
			comic.LastChapter = append(comic.LastChapter, chapter)
		})

		comics = append(comics, comic)
	})

	// Menyiapkan payload yang akan digunakan di post request
	payload := map[string]string{
		"action":               "madara_load_more",
		"page":                 page,
		"template":             "madara-core/content/content-archive",
		"vars[orderby]":        "meta_value_num",
		"vars[posts_per_page]": limit,
		"vars[post_type]":      "wp-manga",
		"vars[meta_key]":       "_latest_update",
	}
	err := c.Post("https://klikmanga.com/wp-admin/admin-ajax.php", payload)
	if err != nil {
		return nil, err
	}
	return comics, nil
}
