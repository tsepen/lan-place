package models

type Site struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	URL         string      `json:"url"`
	Created     string      `json:"createdAt"`
	Rating      int         `json:"rating"`
	Views       int         `json:"views"`
	Tags        []*Tag      `json:"tags"`
	Categories  []*Category `json:"categories"`
}
