package models

// Article represent the article structs
type News struct {
	Id 	  string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}