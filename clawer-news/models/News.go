package models

type News struct {
	Title   string `json:"title"`
	Link    string `json:"link"`
	Src     string `json:"src"`
	PubDate string `json:"pubDate"`
}
