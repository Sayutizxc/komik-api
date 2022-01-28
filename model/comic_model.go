package model

type Comic struct {
	Url         string    `json:"url"`
	Title       string    `json:"title"`
	Thumbnail   string    `json:"thumbnail"`
	Rating      string    `json:"rating"`
	LastChapter []Chapter `json:"last_chapter"`
}
