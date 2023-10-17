package entity

import "time"

type Candidate struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	County    string    `json:"county"`
	Votes     int       `json:"votes"`
}
