package model

type Comic struct {
	Url         string    `json:"url"`
	Title       string    `json:"title"`
	Thumbnail   string    `json:"thumbnail"`
	Score       string    `json:"score"`
	LastChapter []Chapter `json:"last_chapter"`
}
