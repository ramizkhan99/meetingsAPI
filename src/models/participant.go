package models

// Participant : Model for participant
type Participant struct {
	Name 	string	`json:"name"`
	Email 	string	`json:"email"`
	RSVP 	string	`json:"rsvp"`
}