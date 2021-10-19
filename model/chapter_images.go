package model

type ChapterImages struct {
	Chapter  string   `json:"chapter"`
	Url      string   `json:"url"`
	Next     string   `json:"next"`
	Prev     string   `json:"prev"`
	Images   []string `json:"images"`
	Comments []Comment
}
