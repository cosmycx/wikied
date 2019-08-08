package model

import "time"

// Post Api model
type PostIn struct {
	Originator string `json:"originator"`
	Content    string `json:"content"`
} // .Post Api model

// Post backend model
type Post struct {
	// from UI
	Originator string `json:"originator"`
	Content    string `json:"content"`
	// added BE
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
} // .Post
