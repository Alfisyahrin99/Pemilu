package model

import "time"

type Missions struct {
	Id          int
	CandidateId int
	Mission     string
	CreatedAt   time.Time
	UpdatedAt	time.Time
}