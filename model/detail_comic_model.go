package model

type DetailComic struct {
	Url       string    `json:"url"`
	Title     string    `json:"title"`
	Thumbnail string    `json:"thumbnail"`
	Rating    string    `json:"rating"`
	Authors   string    `json:"authors"`
	Artists   string    `json:"artists"`
	Genres    []string  `json:"genres"`
	Type      string    `json:"type"`
	Release   string    `json:"release"`
	Status    string    `json:"status"`
	Synopsis  string    `json:"synopsis"`
	Chapters  []Chapter `json:"chapters"`
	Comments  []Comment `json:"comments"`
}
