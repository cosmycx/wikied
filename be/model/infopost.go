package model

import "time"

// InfoPost backend model
type InfoPost struct {
	// from UI
	Content    			string `json:"content"`
	// added BE
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
} // .InfoPost
