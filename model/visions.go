package model

import "time"

type Visions struct {
	Id          int
	CandidateId int
	Visions     string
	CreatedAt   time.Time
	UpdatedAt	time.Time
}