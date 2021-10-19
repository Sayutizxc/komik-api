package repository

import (
	"github.com/gocolly/colly"
	"github.com/sayutizxc/klikmanga-scraper/model"
	"strings"
)

func SearchComicRepository(s string, page string, limit string) ([]model.Comic, error) {
	comics, err := getSearchComicRepository(s, page, limit)
	if err != nil {
		return nil, err
	}
	return comics, nil
}

func getSearchComicRepository(s string, page string, limit string) ([]model.Comic, error) {
	var comics []model.Comic
	c := colly.NewCollector()

	c.OnHTML(".row", func(e *colly.HTMLElement) {
		var comic model.Comic
		comic.Url = e.ChildAttr("div.tab-summary > div.post-title > h3 > a", "href")
		comic.Title = e.ChildText("div.tab-summary > div.post-title > h3 > a")
		comic.Thumbnail = e.ChildAttr("div > a > img", "src")
		comic.Score = e.ChildText("div.tab-meta > div.meta-item.rating > div > span")

		var chapter model.Chapter
		chapter.ChapterUrl = e.ChildAttr("div.tab-meta > div.meta-item.latest-chap > span.font-meta.chapter > a", "href")
		chapter.Chapter = e.ChildText("div.tab-meta > div.meta-item.latest-chap > span.font-meta.chapter > a")
		chapter.Date = e.ChildText("div.tab-meta > div.meta-item.post-on > span")
		if chapter.Date == "" {
			chapter.Date = e.ChildAttr("div.tab-meta > div.meta-item.post-on > span > span > a", "title")
			chapter.Date = strings.Replace(chapter.Date, "ago", "yang lalu", 1)
		}

		comic.LastChapter = append(comic.LastChapter, chapter)
		comics = append(comics, comic)
	})

	payload := map[string]string{
		"action":               "madara_load_more",
		"page":                 page,
		"template":             "madara-core/content/content-search",
		"vars[posts_per_page]": limit,
		"vars[post_type]":      "wp-manga",
		"vars[s]":              s,
	}
	err := c.Post("https://klikmanga.com/wp-admin/admin-ajax.php", payload)
	if err != nil {
		return nil, err
	}
	return comics, nil
}
