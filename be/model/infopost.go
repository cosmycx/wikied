package model

import "time"

// InfoPost backend model
type InfoPost struct {
	// from UI
	Originator 			string `json:"originator"`
	Content    			string `json:"content"`
	ContentType       	string `json:"content_type"`
	// added BE
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
} // .InfoPost
