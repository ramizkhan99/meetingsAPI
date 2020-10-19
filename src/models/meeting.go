package models

import	"time"

// Meeting : Model for meeting
type Meeting struct {
	Title 			string 			`json:"title"`
	Participants 	[]Participant 	`json:"participants"`
	StartAt 		time.Time 		`json:"start"`
	EndAt 			time.Time 		`json:"end"`
	CreatedAt 		time.Time 		`json:"created"`
}